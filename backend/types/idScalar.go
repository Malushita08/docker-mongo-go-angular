package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var IdScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "IdScalarType",
	Description: "The `IdScalarType` scalar type represents an ID Object.",
	// Serialize serializes `CustomID` to string.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			v := *value
			return v.Hex()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `CustomID`.
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			oid, err := primitive.ObjectIDFromHex(value)
			if err!=nil{
				return nil
			}
			return oid
		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST value to `CustomID`.
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			oid, err := primitive.ObjectIDFromHex(valueAST.Value)
			if err!=nil{
				return nil
			}
			return oid
		default:
			return nil
		}
	},
})
