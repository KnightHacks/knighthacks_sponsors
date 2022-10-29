package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KnightHacks/knighthacks_shared/pagination"
	"github.com/KnightHacks/knighthacks_sponsors/graph/generated"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
)

func (r *mutationResolver) CreateSponsor(ctx context.Context, input model.NewSponsor) (*model.Sponsor, error) {
	return r.Repository.CreateSponsor(ctx, &input)
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, id string, input model.UpdatedSponsor) (*model.Sponsor, error) {
	if input.Description == nil && input.Name == nil && input.Logo == nil && input.Tier == nil && input.Website == nil && input.Since == nil {
		return nil, fmt.Errorf("no field has been updated")
	}

	return r.Repository.UpdateSponsor(ctx, id, &input)
}

func (r *mutationResolver) DeleteSponsor(ctx context.Context, id string) (bool, error) {
	return r.Repository.DeleteSponsor(ctx, id)
}

func (r *queryResolver) Sponsors(ctx context.Context, filter *model.SponsorFilter, first int, after *string) (*model.SponsorsConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	sponsors, total, err := r.Repository.GetSponsors(ctx, first, a)
	if err != nil {
		return nil, err
	}

	return &model.SponsorsConnection{
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(sponsors[0].ID, sponsors[len(sponsors)-1].ID),
		Sponsors:   sponsors,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
