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
	//TODO implement me
	panic("implement me")
}
