package database

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var (
	DB_USER ="admin"
	DB_PASSWORD ="101250"
	DB_HOST ="172.17.0.3"
	DB_PORT ="27017"
	DB_NAME ="db_books"

	//MongoDBHosts = fmt.Sprintf("mongodb://%s:%s@%s:%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT)
	MongoDBHosts = fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	//MongoDBHosts = fmt.Sprintf("mongodb://%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	AuthDatabase = os.Getenv("DB_NAME")
	//AuthDatabase = DB_NAME
)

var (
	Session *mongo.Client
	DBError error
	Ctx     context.Context
)

var (
	COLLECTIONS_NAME = []string{
		"users",
		"modules",
		"operations",
		"sections",
		"contractualDefects",
		"pcis",
		"photos",
		"sectors",
	}
)

func ConnectDB() (*mongo.Client, context.Context, error) {
	//Set up a context required by mongo.Connect
	Ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	//Create connection
	Session, DBError = mongo.NewClient(options.Client().ApplyURI(MongoDBHosts))
	if DBError != nil {
		return nil, nil, DBError
	}
	//Connect
	DBError = Session.Connect(Ctx)
	if DBError != nil {
		return nil, nil, DBError
	}

	//fmt.Println("conexion", MongoDBHosts)
	//fmt.Println("db", DBError)
	return Session, Ctx, DBError
}

func DropAndInitDataBase() bool {
	if Session == nil {
		return false
	}
	sessionCopy := Session
	err := sessionCopy.Database(AuthDatabase).Drop(Ctx)
	if err != nil {
		return false
	}
	c, err := GetCollection("installs")
	if err != nil {
		return false
	}
	m := primitive.M{
		"date": time.Now(),
	}
	_, err = c.InsertOne(Ctx, m)
	if err != nil {
		return false
	}
	collections, err := sessionCopy.Database(AuthDatabase).ListCollectionNames(Ctx, bson.M{})
	if err != nil {
		return false
	}

	for _, col := range COLLECTIONS_NAME {
		exists := false
		for _, col_ := range collections {
			if col == col_ {
				exists = true
				break
			}
		}
		if !exists {
			_, err := sessionCopy.Database(AuthDatabase).Collection(col).InsertOne(Ctx, bson.M{})
			if err != nil {
				return false
			}
			_, err = sessionCopy.Database(AuthDatabase).Collection(col).DeleteMany(Ctx, bson.M{})
			if err != nil {
				return false
			}
		}
	}
	return true
}

func GetCollection(collection string) (*mongo.Collection, error) {
	if Session == nil {
		return nil, errors.New("error Session Empty")
	}
	return Session.Database(AuthDatabase).Collection(collection), nil
}

func CloseSession(session *mongo.Client) {
	if session != nil {
		_ = session.Disconnect(Ctx)
	}
}
