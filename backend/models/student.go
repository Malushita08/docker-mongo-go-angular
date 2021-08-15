package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	LastName string             `bson:"lastName" json:"lastName"`
	DNI      string             `bson:"dni" json:"dni"`
	Code     string             `bson:"code" json:"code"`
}

func StudentAddToBson(studentI interface{}) (*primitive.ObjectID, *bson.M, error) {
	studentM, ok := studentI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting student")
	}
	id := primitive.NewObjectID()
	name, _ := studentM["name"].(string)
	lastName, _ := studentM["lastName"].(string)
	dni, _ := studentM["dni"].(string)
	code, _ := studentM["code"].(string)
	studentBson := bson.M{
		"_id":      id,
		"name":     name,
		"lastName": lastName,
		"dni":      dni,
		"code":     code,
	}
	return &id, &studentBson, nil
}

func StudentUpdateToBson(idW interface{}, studentI interface{}) (*primitive.ObjectID, *bson.M, error) {
	idwM, ok := idW.(string)
	if !ok {
		return nil, nil, errors.New("error casting id student")
	}
	idwm, _ := primitive.ObjectIDFromHex(idwM)
	studentM, ok := studentI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting student")
	}
	name, _ := studentM["name"].(string)
	lastName, _ := studentM["lastName"].(string)
	dni, _ := studentM["dni"].(string)
	code, _ := studentM["code"].(string)
	studentBson := bson.M{
		"name":     name,
		"lastName": lastName,
		"dni":      dni,
		"code":     code,
	}
	return &idwm, &studentBson, nil
}
