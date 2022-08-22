package repository

import (
	"context"
	"strconv"

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
	var sponsorId string
	var sponsorIdInt int
	err := r.DatabasePool.QueryRow(ctx, "INSERT INTO sponsors (name, tier, since, description, website, logo) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		input.Name,
		input.Tier,
		input.Since,
		input.Description,
		input.Website,
		input.Logo,
	).Scan(&sponsorIdInt)
	if err != nil {
		return nil, err
	}
	sponsorId = strconv.Itoa(sponsorIdInt)

	return &model.Sponsor {
		ID: 			sponsorId,
		Name: 			input.Name,
		Tier: 			input.Tier,
		Since: 			*input.Since,
		Description: 	input.Description,
		Website: 		input.Website,
		Logo: 			input.Logo,
	}, nil
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
