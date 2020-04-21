package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/242617/synapse-core/db/model"
	"github.com/242617/synapse-core/graphql/generated"
)

func (r *mutationResolver) CreateCrawler(ctx context.Context, input model.NewCrawler) (*model.Crawler, error) {
	id, _, err := r.db.CreateCrawler(input.Name, input.Certificate)
	if err != nil {
		return nil, err
	}
	return &model.Crawler{
		ID:          id,
		Name:        input.Name,
		Certificate: input.Certificate,
	}, nil
}

func (r *mutationResolver) DeleteCrawler(ctx context.Context, id int) (*model.Crawler, error) {
	crawler, err := r.db.GetCrawler(id)
	if err != nil {
		return nil, err
	}
	if err := r.db.DeleteCrawler(id); err != nil {
		return nil, err
	}
	return crawler, nil
}

func (r *queryResolver) Crawlers(ctx context.Context, id *int) ([]*model.Crawler, error) {
	if id != nil {
		crawler, err := r.db.GetCrawler(*id)
		if err != nil {
			return nil, err
		}
		return []*model.Crawler{crawler}, nil
	}
	crawlers, err := r.db.GetCrawlers()
	if err != nil {
		return nil, err
	}
	return crawlers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
