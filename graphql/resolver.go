package graphql

import "github.com/242617/synapse-core/db"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{ db db.DB }

func NewResolver(database db.DB) *Resolver {
	return &Resolver{db: database}
}
