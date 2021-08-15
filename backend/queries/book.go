package queries

import (
	"github.com/graphql-go/graphql"
	"backend-crud/services"
	"backend-crud/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var GetBookQuery = &graphql.Field{
	Type:        types.BookType,
	Description: "Get Book by Id",
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
		Book, err := services.ReadBook(id)
		if err != nil {
			return nil, err
		}
		return *Book, nil
	},
}

var GetBooksQuery = &graphql.Field{
	Description: "Get all packages",
	Type:        graphql.NewList(types.BookType),
	Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
		packages, err := services.ReadBooks()
		if err != nil {
			return nil, err
		}
		return packages, nil
	},
}
