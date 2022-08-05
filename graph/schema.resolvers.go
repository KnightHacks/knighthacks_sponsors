package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KnightHacks/knighthacks_sponsors/graph/generated"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
)

func (r *mutationResolver) CreateSponsor(ctx context.Context, input model.NewSponsor) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, id string, input model.UpdatedSponsor) (*model.Sponsor, error) {
	//panic(fmt.Errorf("not implemented"))
	if input.Description == nil && input.Name == nil && input.Logo == nil && input.Tier == nil && input.Website == nil && input.Since == nil {
		return nil, fmt.Errorf("no field has been updated")
	}

	return r.Repository.UpdateSponsor(ctx, id, &input)
}

func (r *mutationResolver) DeleteSponsor(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sponsors(ctx context.Context, filter *model.SponsorFilter) ([]*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
