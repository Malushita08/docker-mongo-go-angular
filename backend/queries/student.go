package queries

import (
	"github.com/graphql-go/graphql"
	"backend-crud/services"
	"backend-crud/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var GetStudentQuery = &graphql.Field{
	Type:        types.StudentType,
	Description: "Get Student by Id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
		if err != nil {
			return nil, err
		}
		Student, err := services.ReadStudent(id)
		if err != nil {
			return nil, err
		}
		return *Student, nil
	},
}

var GetStudentsQuery = &graphql.Field{
	Description: "Get all packages",
	Type:        graphql.NewList(types.StudentType),
	Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
		packages, err := services.ReadStudents()
		if err != nil {
			return nil, err
		}
		return packages, nil
	},
}
