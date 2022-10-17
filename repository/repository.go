package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
)

type Repository interface {
	CreateSponsor(ctx context.Context, input *model.NewSponsor) (*model.Sponsor, error)
	UpdateSponsor(ctx context.Context, id string, input *model.UpdatedSponsor) (*model.Sponsor, error)
	DeleteSponsor(ctx context.Context, id string) (bool, error)
	Sponsors(ctx context.Context)
	GetSponsor(ctx context.Context, id string) (*model.Sponsor, error)
}
