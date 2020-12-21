package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	/* GraphQL */
	graphQL := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Use(middleware.Logger)
	r.Handle("/query", graphQL)
	return r
}

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
