package repository

type Repository interface {
	// TODO: create interface functions & implement them in database_repository.go
	CreateSponsor(ctx context.Context, input *model.NewSponsor) (*model.Sponsor, error)
	UpdateSponsor(ctx context.Context, id string, input *model.UpdatedSponsor) (*model.Sponsor, error)
	DeleteSponsor(ctx context.Context, id string) (*model.Sponsor, error)
	GetSponsor(ctx context.Context, id string) (*model.Sponsor, error)
}
