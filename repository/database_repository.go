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

//DatabaseRepository
//Implements the Repository interface's functions
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

func (r *DatabaseRepository) getSponsorWithQueryable(ctx context.Context, id string, queryable database.Queryable) (*model.Sponsor, error) {
	var event model.Sponsor
	err := queryable.QueryRow(ctx, "SELECT id, description, name, logo, tier, website, since FROM events WHERE id = $1", id).Scan(&event.ID, &event.Description,
		&event.Name, &event.Logo, &event.Tier, &event.Website, &event.Since)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, SponsorNotFound
		}
		return nil, err
	}

	return &event, err
}

func (r *DatabaseRepository) UpdateSponsor(ctx context.Context, id string, input *model.UpdatedSponsor) (*model.Sponsor, error) {
	if input.Description == nil && input.Name == nil && input.Logo == nil && input.Tier == nil && input.Website == nil && input.Since == nil {
		return nil, fmt.Errorf("empty sponsor field")
	}

	var event *model.Sponsor
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
		event, err = r.getSponsorWithQueryable(ctx, id, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *DatabaseRepository) UpdateDesc(ctx context.Context, id string, sponsorDesc string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET description = $1 WHERE id = $2", sponsorDesc, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateName(ctx context.Context, id string, sponsorName string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET name = $1 WHERE id = $2", sponsorName, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateLogo(ctx context.Context, id string, sponsorLogo string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET logo_url = $1 WHERE id = $2", sponsorLogo, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateTier(ctx context.Context, id string, sponsorTier model.SubscriptionTier, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET tier = $1 WHERE id = $2", sponsorTier, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateWebsite(ctx context.Context, id string, sponsorSite string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET website = $1 WHERE id = $2", sponsorSite, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateSince(ctx context.Context, id string, sponsorSince time.Time, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET since = $1 WHERE id = $2", sponsorSince, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return SponsorNotFound
	}
	return nil
}

func (r *DatabaseRepository) DeleteSponsor(ctx context.Context, id string) (*model.Sponsor, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetSponsor(ctx context.Context, id string) (*model.Sponsor, error) {
	//TODO implement me
	panic("implement me")
}
