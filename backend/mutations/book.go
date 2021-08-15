package mutations

import (
	"github.com/graphql-go/graphql"
	"backend-crud/models"
	"backend-crud/services"
	"backend-crud/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddBookMutation = &graphql.Field{
	Type:        types.BookType,
	Description: "Add book",
	Args: graphql.FieldConfigArgument{
		"book": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.BookAddInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//convert book to bson
		Book, bookBson, err := models.BookAddToBson(params.Args["book"])
		if err != nil {
			return nil, err
		}
		//add book service
		if err := services.CreateBook(*bookBson); err != nil {
			return nil, err
		}
		bookResult, err := services.ReadBook(*Book)
		if err != nil {
			return nil, err
		}
		return *bookResult, nil
	},
}

var UpdateBookMutation = &graphql.Field{
	Type:        types.BookType,
	Description: "Update book by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"book": &graphql.ArgumentConfig{
			Type: types.BookUpdateInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//convert book to Bson
		Book, bookBson, err := models.BookUpdateToBson(params.Args["id"], params.Args["book"])
		if err != nil {
			return nil, err
		}
		//update book service
		err = services.UpdateBook(*Book, *bookBson)
		if err != nil {
			return nil, err
		}
		//return book updated
		bookResult, err := services.ReadBook(*Book)
		if err != nil {
			return nil, err
		}
		return *bookResult, nil
	},
}

var DeleteBookMutation = &graphql.Field{
	Type:        types.BookType,
	Description: "Delete book by id",
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
		Book, _ := services.ReadBook(id)
		_, err2 := services.DeleteBook(id)

		if err2 != nil {
			return nil, err
		}
		return *Book, nil
	},
}
