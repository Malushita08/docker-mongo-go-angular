package types

import (
	"github.com/graphql-go/graphql"
)

var StudentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: IdScalarType,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"dni": &graphql.Field{
				Type: graphql.String,
			},
			"code": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var StudentAddInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "StudentAddInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"dni": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var StudentUpdateInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "StudentUpdateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"code": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"dni": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
