package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/djay-S/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://demo:lYIXqRCa9N9pzOICxmWP@cluster0.nevdgmh.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// Most Important
var collection *mongo.Collection

// connect with mongoDB
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection reference
	fmt.Println("Collection reference is ready")

}

// MongoDB helpers - file

// insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in DB with Id:", inserted.InsertedID)
}
