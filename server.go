package main

import (
	"log"
	"net/http"
	"os"

	"example.com/subgraph-template-go-gqlgen-boilerplate/graph"
	"example.com/subgraph-template-go-gqlgen-boilerplate/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://studio.apollographql.com"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", srv)

	log.Printf("Explore with \"https://studio.apollographql.com/sandbox/explorer?endpoint=http://localhost:" + port + "\"")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
