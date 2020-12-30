package types

import (
	"github.com/graphql-go/graphql"
)

var Todo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: ID,
			},
			"task": &graphql.Field{
				Type: graphql.String,
			},
			"is_done": &graphql.Field{
				Type: graphql.Boolean,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"deadline": &graphql.Field{
				Type: graphql.DateTime,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updated_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var ViewListTodoMe = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ViewListTodoMe",
		Fields: graphql.Fields{
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: graphql.NewList(Todo),
			},
		},
	},
)
