package services

import (
	"backend-crud/database"
	"backend-crud/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colBooks = "books"

func CreateBook(book bson.M) error {
	c, errC := database.GetCollection(colBooks)
	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(database.Ctx, book)
	if errC != nil {
		return errC
	}
	return nil
}

func ReadBook(bookId primitive.ObjectID) (*models.Book, error) {
	c, errC := database.GetCollection(colBooks)
	if errC != nil {
		return nil, errC
	}
	var result models.Book
	errC = c.FindOne(database.Ctx, bson.M{"_id": bookId}).Decode(&result)
	return &result, errC
}

func ReadBooks() ([]models.Book, error) {
	c, errC := database.GetCollection(colBooks)
	if errC != nil {
		return nil, errC
	}
	var results []models.Book
	cursor, err := c.Find(database.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(database.Ctx, &results)
	return results, err
}

func UpdateBook(bookId primitive.ObjectID, book bson.M) error {
	c, errorC := database.GetCollection(colBooks)
	if errorC != nil {
		return errorC
	}
	result, err := c.UpdateOne(database.Ctx, bson.M{"_id": bookId}, bson.M{"$set": book})
	if err != nil {
		return err
	}
	if result.MatchedCount != int64(1) {
		return fmt.Errorf("Error not Found. %v ")
	}
	return nil
}

func DeleteBook(bookId primitive.ObjectID) (bool, error) {
	c, errC := database.GetCollection(colBooks)
	if errC != nil {
		return false, errC
	}
	result, err := c.DeleteOne(database.Ctx, bson.M{"_id": bookId})
	if err != nil {
		return false, err
	}
	if result.DeletedCount != 1 {
		return false, fmt.Errorf("Error not Found. %v ")
	}
	return true, err
}
