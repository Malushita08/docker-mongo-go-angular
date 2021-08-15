package mutations

import (
	"github.com/graphql-go/graphql"
	"backend-crud/models"
	"backend-crud/services"
	"backend-crud/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddStudentMutation = &graphql.Field{
	Type:        types.StudentType,
	Description: "Add student",
	Args: graphql.FieldConfigArgument{
		"student": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.StudentAddInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//convert student to bson
		Student, studentBson, err := models.StudentAddToBson(params.Args["student"])
		if err != nil {
			return nil, err
		}
		//add student service
		if err := services.CreateStudent(*studentBson); err != nil {
			return nil, err
		}
		studentResult, err := services.ReadStudent(*Student)
		if err != nil {
			return nil, err
		}
		return *studentResult, nil
	},
}

var UpdateStudentMutation = &graphql.Field{
	Type:        types.StudentType,
	Description: "Update student by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"student": &graphql.ArgumentConfig{
			Type: types.StudentUpdateInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//convert student to Bson
		Student, studentBson, err := models.StudentUpdateToBson(params.Args["id"], params.Args["student"])
		if err != nil {
			return nil, err
		}
		//update student service
		err = services.UpdateStudent(*Student, *studentBson)
		if err != nil {
			return nil, err
		}
		//return student updated
		studentResult, err := services.ReadStudent(*Student)
		if err != nil {
			return nil, err
		}
		return *studentResult, nil
	},
}

var DeleteStudentMutation = &graphql.Field{
	Type:        types.StudentType,
	Description: "Delete student by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, err := primitive.ObjectIDFromHex(params.Args["id"].(string))
		if err != nil {
			return nil, err
		}
		Student, _ := services.ReadStudent(id)
		_, err2 := services.DeleteStudent(id)

		if err2 != nil {
			return nil, err
		}
		return *Student, nil
	},
}
