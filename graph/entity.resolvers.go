package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/KnightHacks/knighthacks_sponsors/graph/generated"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
)

func (r *entityResolver) FindSponsorByID(ctx context.Context, id string) (*model.Sponsor, error) {
	return &model.Sponsor{ID: id}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
