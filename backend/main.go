package main

import (
	"backend-crud/database"
	graph "backend-crud/graphql"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

//func main() {
//	fmt.Println("Conexión a MongoDB")
//}

//var schema, _ = graphql.NewSchema(
//	graphql.SchemaConfig{
//		Query:    ,
//		Mutation: mutationType,
//	},
//)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w,r)
	})
}

func main() {
	_, _, err := database.ConnectDB()
	if err != nil {
		fmt.Errorf("Error db: %v", err)
	}
	schema := *graph.GetSchema()

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	//http.Handle("/graphql", h)
	http.Handle("/graphql", CorsMiddleware(h))
	//
	//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	//	result := executeQuery(r.URL.Query().Get("query"), schema)
	//	json.NewEncoder(w).Encode(result)
	//})

	fmt.Println("Server is running on port 8001")
	//http.ListenAndServe(":8080", nil)

	log.Fatalln(http.ListenAndServe(":8001", nil))
	//}
}

//package main
//

//
////func main() {
////	fmt.Println("Conexión a MongoDB")
////}
//
////var schema, _ = graphql.NewSchema(
////	graphql.SchemaConfig{
////		Query:    ,
////		Mutation: mutationType,
////	},
////)
//
//func executeQuery(query string, schema graphql.Schema) *graphql.Result {
//	result := graphql.Do(graphql.Params{
//		Schema:        schema,
//		RequestString: query,
//	})
//	if len(result.Errors) > 0 {
//		fmt.Printf("errors: %v", result.Errors)
//	}
//	return result
//}
//
//func main() {
//	_, _, err := database.ConnectDB()
//	if err != nil {
//		fmt.Println(err)
//	}
//	schema := *graph.GetSchema()
//
//	h := handler.New(&handler.Config{
//		Schema:     &schema,
//		Pretty:     true,
//		GraphiQL:   false,
//		Playground: true,
//	})
//
//	http.Handle("/graphql", h)
//
//	//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
//	//	result := executeQuery(r.URL.Query().Get("query"), schema)
//	//	json.NewEncoder(w).Encode(result)
//	//})
//
//	log.Printf("connect to http://localhost for GraphQL playground 8080")
//	//http.ListenAndServe(":8080", nil)
//
//	log.Fatal(http.ListenAndServe(":8001", nil))
//	//}
//}
