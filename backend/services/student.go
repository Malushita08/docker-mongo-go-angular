package services

import (
	"backend-crud/database"
	"backend-crud/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colStudents = "students"

func CreateStudent(student bson.M) error {
	c, errC := database.GetCollection(colStudents)
	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(database.Ctx, student)
	if errC != nil {
		return errC
	}
	return nil
}

func ReadStudent(studentId primitive.ObjectID) (*models.Student, error) {
	c, errC := database.GetCollection(colStudents)
	if errC != nil {
		return nil, errC
	}
	var result models.Student
	errC = c.FindOne(database.Ctx, bson.M{"_id": studentId}).Decode(&result)
	return &result, errC
}

func ReadStudents() ([]models.Student, error) {
	c, errC := database.GetCollection(colStudents)
	if errC != nil {
		return nil, errC
	}
	var results []models.Student
	cursor, err := c.Find(database.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(database.Ctx, &results)
	return results, err
}

func UpdateStudent(studentId primitive.ObjectID, student bson.M) error {
	c, errorC := database.GetCollection(colStudents)
	if errorC != nil {
		return errorC
	}
	result, err := c.UpdateOne(database.Ctx, bson.M{"_id": studentId}, bson.M{"$set": student})
	if err != nil {
		return err
	}
	if result.MatchedCount != int64(1) {
		return fmt.Errorf("Error not Found. %v ")
	}
	return nil
}

func DeleteStudent(studentId primitive.ObjectID) (bool, error) {
	c, errC := database.GetCollection(colStudents)
	if errC != nil {
		return false, errC
	}
	result, err := c.DeleteOne(database.Ctx, bson.M{"_id": studentId})
	if err != nil {
		return false, err
	}
	if result.DeletedCount != 1 {
		return false, fmt.Errorf("Error not Found. %v ")
	}
	return true, err
}
