package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"time"
)

var (
	SponsorNotFound = errors.New("sponsor was not found")
)

// DatabaseRepository
// Implements the Repository interface's functions
type DatabaseRepository struct {
	DatabasePool *pgxpool.Pool
}

func NewDatabaseRepository(databasePool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePool: databasePool,
	}
}

func (r *DatabaseRepository) CreateSponsor(ctx context.Context, input *model.NewSponsor) (*model.Sponsor, error) {
	var sponsorIdInt int
	err := r.DatabasePool.QueryRow(ctx, "INSERT INTO sponsors (name, tier, since, description, website, logo_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		&input.Name,
		&input.Tier,
		&input.Since,
		&input.Description,
		&input.Website,
		&input.Logo,
	).Scan(&sponsorIdInt)
	if err != nil {
		return nil, err
	}

	return &model.Sponsor{
		ID:          strconv.Itoa(sponsorIdInt),
		Name:        input.Name,
		Tier:        input.Tier,
		Since:       *input.Since,
		Description: input.Description,
		Website:     input.Website,
		Logo:        input.Logo,
	}, nil
}

func (r *DatabaseRepository) GetSponsorWithQueryable(ctx context.Context, id string, queryable database.Queryable) (*model.Sponsor, error) {
	var sponsor model.Sponsor
	var idInt int

	err := queryable.QueryRow(ctx, "SELECT id, description, name, logo_url, tier, website, since FROM sponsors WHERE id = $1", id).Scan(
		&idInt,
		&sponsor.Description,
		&sponsor.Name,
		&sponsor.Logo,
		&sponsor.Tier,
		&sponsor.Website,
		&sponsor.Since,
	)

	sponsor.ID = strconv.Itoa(idInt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, SponsorNotFound
		}
		return nil, err
	}

	return &sponsor, err
}

func (r *DatabaseRepository) UpdateSponsor(ctx context.Context, id string, input *model.UpdatedSponsor) (*model.Sponsor, error) {
	if input.Description == nil && input.Name == nil && input.Logo == nil && input.Tier == nil && input.Website == nil && input.Since == nil {
		return nil, fmt.Errorf("empty sponsor field")
	}

	var sponsor *model.Sponsor
	var err error
	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if input.Description != nil {
			err = r.UpdateDesc(ctx, id, *input.Description, tx)
			if err != nil {
				return err
			}
		}
		if input.Name != nil {
			err = r.UpdateName(ctx, id, *input.Name, tx)
			if err != nil {
				return err
			}
		}
		if input.Logo != nil {
			err = r.UpdateLogo(ctx, id, *input.Logo, tx)
			if err != nil {
				return err
			}
		}
		if input.Tier != nil {
			err := r.UpdateTier(ctx, id, *input.Tier, tx)
			if err != nil {
				return err
			}
		}
		if input.Website != nil {
			err := r.UpdateWebsite(ctx, id, *input.Website, tx)
			if err != nil {
				return err
			}
		}
		if input.Since != nil {
			err := r.UpdateSince(ctx, id, *input.Since, tx)
			if err != nil {
				return err
			}
		}
		sponsor, err = r.GetSponsorWithQueryable(ctx, id, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return sponsor, nil
}

func (r *DatabaseRepository) UpdateDesc(ctx context.Context, id string, sponsorDesc string, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET description = $1 WHERE id = $2", sponsorDesc, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateName(ctx context.Context, id string, sponsorName string, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET name = $1 WHERE id = $2", sponsorName, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateLogo(ctx context.Context, id string, sponsorLogo string, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET logo_url = $1 WHERE id = $2", sponsorLogo, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateTier(ctx context.Context, id string, sponsorTier model.SubscriptionTier, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET tier = $1 WHERE id = $2", sponsorTier, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateWebsite(ctx context.Context, id string, sponsorSite string, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET website = $1 WHERE id = $2", sponsorSite, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateSince(ctx context.Context, id string, sponsorSince time.Time, queryable database.Queryable) error {
	commandTag, err := queryable.Exec(ctx, "UPDATE sponsors SET since = $1 WHERE id = $2", sponsorSince, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) DeleteSponsor(ctx context.Context, id string) (bool, error) {

	// removes sponsors
	commandTag, err := r.DatabasePool.Exec(ctx, "DELETE FROM sponsors WHERE id = $1", id)

	// checks if there is an error
	if err != nil {
		return false, err
	}
	// checking to see if there is 1 row affected for deleted sponsors if not there is an issue
	if commandTag.RowsAffected() != 1 {
		return false, SponsorNotFound
	}

	// if the above conditions dont execute everything is good
	return true, nil
}

func (r *DatabaseRepository) GetSponsor(ctx context.Context, id string) (*model.Sponsor, error) {
	return r.GetSponsorWithQueryable(ctx, id, r.DatabasePool)
}

func (r *DatabaseRepository) GetSponsors(ctx context.Context, filter *model.SponsorFilter, first int, after string) (sponsors []*model.Sponsor, total int, err error) {
	var sql string
	var variables []any
	var totalSql string
	var totalVariables []any

	if filter != nil {
		stringTiers := make([]string, 0, len(filter.Tiers))

		for _, tier := range filter.Tiers {
			stringTiers = append(stringTiers, tier.String())
		}

		sql = `SELECT id, description, name, logo_url, tier, website, since FROM sponsors WHERE tier = ANY ($1) AND id > $2 LIMIT $3`
		variables = []any{stringTiers, after, first}

		totalSql = `SELECT COUNT(*) FROM sponsors WHERE tier = ANY ($1) AND id > $2`
		totalVariables = []any{stringTiers, after}
	} else {
		sql = `SELECT id, description, name, logo_url, tier, website, since FROM sponsors WHERE id > $1 LIMIT $2`
		variables = []any{after, first}

		totalSql = `SELECT COUNT(*) FROM sponsors WHERE id > $1`
		totalVariables = []any{after}
	}
	err = pgx.BeginTxFunc(ctx, r.DatabasePool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if err := tx.QueryRow(ctx, totalSql, totalVariables...).Scan(&total); err != nil {
			return err
		}

		if total > 0 {
			rows, err := tx.Query(ctx, sql, variables...)
			if err != nil {
				return err
			}

			for rows.Next() {
				var sponsor model.Sponsor
				var intId int
				err = rows.Scan(
					&intId,
					&sponsor.Description,
					&sponsor.Name,
					&sponsor.Logo,
					&sponsor.Tier,
					&sponsor.Website,
					&sponsor.Since,
				)
				if err != nil {
					return err
				}
				sponsor.ID = strconv.Itoa(intId)
				sponsors = append(sponsors, &sponsor)
			}
		}
		return nil
	})

	return sponsors, total, err
}
