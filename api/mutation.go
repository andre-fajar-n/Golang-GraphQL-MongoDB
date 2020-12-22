package api

import (
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/resolver"
	"github.com/andre-fajar-n/Golang-GraphQL-MongoDB/types"
	"github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"register": &graphql.Field{
			Type:        types.User,
			Description: "User Register",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.Register,
		},
		"login": &graphql.Field{
			Type:        types.User,
			Description: "User Login",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.Login,
		},
		"addTodo": &graphql.Field{
			Type:        types.Todo,
			Description: "Add Todo",
			Args: graphql.FieldConfigArgument{
				"task": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"deadline": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.AddTodo,
		},
	},
})
