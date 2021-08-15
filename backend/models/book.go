package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Author string             `bson:"author" json:"author"`
	Tittle string             `bson:"tittle" json:"tittle"`
	Price  float64            `bson:"price" json:"price"`
}

func BookAddToBson(bookI interface{}) (*primitive.ObjectID, *bson.M, error) {
	bookM, ok := bookI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting book")
	}
	id := primitive.NewObjectID()
	author, _ := bookM["author"].(string)
	tittle, _ := bookM["tittle"].(string)
	price, _ := bookM["price"].(float64)
	bookBson := bson.M{
		"_id":    id,
		"author": author,
		"tittle": tittle,
		"price":  price,
	}
	return &id, &bookBson, nil
}

func BookUpdateToBson(idW interface{}, bookI interface{}) (*primitive.ObjectID, *bson.M, error) {
	idwM, ok := idW.(string)
	if !ok {
		return nil, nil, errors.New("error casting id book")
	}
	idwm, _ := primitive.ObjectIDFromHex(idwM)
	bookM, ok := bookI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting book")
	}
	author, _ := bookM["author"].(string)
	tittle, _ := bookM["tittle"].(string)
	price, _ := bookM["price"].(float64)
	bookBson := bson.M{
		"author":   author,
		"tittle": tittle,
		"price":  price,
	}
	return &idwm, &bookBson, nil
}
