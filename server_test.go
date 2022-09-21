package main

import (
	"testing"

	"example.com/subgraph-template-go-gqlgen-boilerplate/graph"
	"example.com/subgraph-template-go-gqlgen-boilerplate/graph/generated"
	"example.com/subgraph-template-go-gqlgen-boilerplate/graph/model"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/stretchr/testify/assert"
)

func TestFooQuery(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	client := client.New(srv)

	fooQuery := `
	query FooQuery {
		foo(id: "1") {
			id
			name
		}
	}
  `
	var resp struct {
		Foo *model.Foo
	}

	client.MustPost(fooQuery, &resp)

	assert.NotNil(t, resp.Foo, "Foo should not be nil")
	assert.Equal(t, "1", resp.Foo.ID, "Foo ID should be the same")
	assert.Equal(t, "Name", *resp.Foo.Name, "Foo Name should be the same")
}
