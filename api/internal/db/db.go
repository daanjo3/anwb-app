package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TODO update connection structure following:
// https://blog.stackademic.com/building-a-robust-rest-api-with-golang-gin-and-mongodb-701faa8961da
func SetupMongoDB() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}
	collection := client.Database("anwb-app").Collection("road-data")
	return collection, client, ctx, cancel
}

func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}

func InsertDocument(data anwb.Document) (primitive.ObjectID, error) {
	collection, client, context, cancel := SetupMongoDB()
	defer CloseConnection(client, context, cancel)

	data.Id = primitive.NewObjectID() // necessary?
	result, err := collection.InsertOne(context, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func Get_Documents() ([]anwb.Document, error) {
	collection, _, context, cancel := SetupMongoDB()
	defer cancel()
	var entryList []anwb.Document
	cursor, err := collection.Find(context, bson.D{})
	defer cursor.Close(context)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context) {
		var entry anwb.Document
		err := cursor.Decode(&entry)
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}
	return entryList, nil
}

func Get_Document_ById(id string) (anwb.Document, error) {
	collection, _, context, cancel := SetupMongoDB()
	defer cancel()
	result := collection.FindOne(context, bson.D{})
	if result.Err() != nil {
		return anwb.Document{}, result.Err()
	}
	var entry anwb.Document
	err := result.Decode(&entry)
	if err != nil {
		return anwb.Document{}, err
	}
	return entry, nil
}
