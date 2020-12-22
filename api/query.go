package api

import (
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/resolver"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/types"
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"viewOneTodoByID": &graphql.Field{
				Type:        types.Todo,
				Description: "View One Todo By ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.ViewOneTodoByID,
			},
		},
	})
