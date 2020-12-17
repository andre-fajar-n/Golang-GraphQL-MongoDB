package app

import "github.com/graphql-go/graphql"

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Query",
		Fields: graphql.Fields{},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: graphql.Fields{},
})

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
