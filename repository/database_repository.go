package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
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

func (r *DatabaseRepository) UpdateSponsor(ctx context.Context, id string, input *model.UpdatedSponsor) (*model.Sponsor, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) DeleteSponsor(ctx context.Context, id string) (*model.Sponsor, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetSponsor(ctx context.Context, id string) (*model.Sponsor, error) {
	var sponsor model.Sponsor
	err := r.DatabasePool.QueryRow(
		ctx,
		"SELECT id, name, tier, since, description, website, logo_url FROM sponsors WHERE id = $1",
		id,
	).Scan(&sponsor.ID, &sponsor.Name, &sponsor.Tier, &sponsor.Since, &sponsor.Description, &sponsor.Website, &sponsor.Logo)

	if err != nil {
		return nil, err
	}

	return &sponsor, err 
}