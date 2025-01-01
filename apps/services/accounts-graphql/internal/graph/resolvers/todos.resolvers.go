package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"apps/services/accounts-graphql/internal/graph/models"
	"context"
	"log"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Println("resolvers.Todos")
	return []*models.Todo{
		{ID: "1", Done: false, Text: "Buy milk"},
	}, nil
}
