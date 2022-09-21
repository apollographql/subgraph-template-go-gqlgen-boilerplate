package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"example.com/subgraph-template-go-gqlgen-boilerplate/graph/generated"
	"example.com/subgraph-template-go-gqlgen-boilerplate/graph/model"
)

// FindFooByID is the resolver for the findFooByID field.
func (r *entityResolver) FindFooByID(ctx context.Context, id string) (*model.Foo, error) {
	return FindFoo(id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
