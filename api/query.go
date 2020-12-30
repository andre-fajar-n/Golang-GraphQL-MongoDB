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
			"viewListTodoMe": &graphql.Field{
				Type:        types.ViewListTodoMe,
				Description: "View List Todo Me",
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"per_page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.ViewListTodoMe,
			},
		},
	})
