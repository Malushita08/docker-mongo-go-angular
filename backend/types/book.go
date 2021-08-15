package types

import (
	"github.com/graphql-go/graphql"
)

var BookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: IdScalarType,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"tittle": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var BookAddInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BookAddInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"author": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"tittle": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})

var BookUpdateInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BookUpdateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"author": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"tittle": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})
