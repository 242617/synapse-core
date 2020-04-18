package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/242617/synapse-core/graphql/generated"
	"github.com/242617/synapse-core/graphql/model"
)

func (r *mutationResolver) CreateCrawler(ctx context.Context, input model.NewCrawler) (*model.Crawler, error) {
	crawler := model.Crawler{
		ID:          strconv.Itoa(rand.Intn(10000)),
		Name:        input.Name,
		Certificate: input.Certificate,
	}
	crawlers = append(crawlers, &crawler)
	return &crawler, nil
}

func (r *queryResolver) Crawlers(ctx context.Context) ([]*model.Crawler, error) {
	return crawlers, nil
}

func (r *queryResolver) Labels(ctx context.Context) ([]*model.Label, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
