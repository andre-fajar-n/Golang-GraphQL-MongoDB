package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	/* GraphQL */
	graphQL := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
		// GraphiQL: true,
		Playground: true,
	})
	r.Use(middleware.Logger)
	r.Handle("/query", headerAuthorization(graphQL))
	return r
}

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

// Header Authorization
func headerAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		auth := r.Header.Get("Authorization")
		ctx = context.WithValue(r.Context(), "token", auth)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
