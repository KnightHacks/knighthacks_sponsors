package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetSponsorWithQueryable(ctx context.Context, id string, queryable database.Queryable) (*model.Sponsor, error) {
	var sponsor model.Sponsor
	err := queryable.QueryRow(ctx, "SELECT id, description, name, logo_url, tier, website, since FROM sponsors WHERE id = $1", id).Scan(&sponsor.ID, &sponsor.Description,
		&sponsor.Name, &sponsor.Logo, &sponsor.Tier, &sponsor.Website, &sponsor.Since)

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
	err = r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
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

func (r *DatabaseRepository) UpdateDesc(ctx context.Context, id string, sponsorDesc string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET description = $1 WHERE id = $2", sponsorDesc, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateName(ctx context.Context, id string, sponsorName string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET name = $1 WHERE id = $2", sponsorName, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateLogo(ctx context.Context, id string, sponsorLogo string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET logo_url = $1 WHERE id = $2", sponsorLogo, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateTier(ctx context.Context, id string, sponsorTier model.SubscriptionTier, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET tier = $1 WHERE id = $2", sponsorTier, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateWebsite(ctx context.Context, id string, sponsorSite string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET website = $1 WHERE id = $2", sponsorSite, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateSince(ctx context.Context, id string, sponsorSince time.Time, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE sponsors SET since = $1 WHERE id = $2", sponsorSince, id)
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

func (r *DatabaseRepository) GetSponsors(ctx context.Context, first int, after string) ([]*model.Sponsor, int, error) {
	//TODO implement me
	panic("implement me")
}
