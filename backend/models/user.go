package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User datos del usuario
type Alumnos struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastName"`
}
