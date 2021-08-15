package graphql

import (
	"backend-crud/mutations"
	"backend-crud/queries"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		//books
		"getBook":  queries.GetBookQuery,
		"getBooks": queries.GetBooksQuery,

		//students
		"getStudent":  queries.GetStudentQuery,
		"getStudents": queries.GetStudentsQuery,
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// students
		"addStudent":    mutations.AddStudentMutation,
		"updateStudent": mutations.UpdateStudentMutation,
		"removeStudent": mutations.DeleteStudentMutation,

		// books
		"addBooks":    mutations.AddBookMutation,
		"updateBooks": mutations.UpdateBookMutation,
		"removeBook":  mutations.DeleteBookMutation,
	},
})

func GetSchema() *graphql.Schema {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
		//Subscription: rootSubscription,
	})
	return &schema
}
