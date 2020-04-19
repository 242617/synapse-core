package graphql

import "github.com/242617/synapse-core/graphql/model"

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen
// It serves as dependency injection for your app, add any dependencies you require here.

var crawlers = []*model.Crawler{}

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}
