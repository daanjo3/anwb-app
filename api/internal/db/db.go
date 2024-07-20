package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var anwbCollection *mongo.Collection
var DATABASE_NAME = "anwb-app"
var COLLECTION_NAME = "road-data"

func GetClient() *mongo.Client {
	if client != nil {
		return client
	}
	uri := os.Getenv("MONGO_CONN")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetAnwbCollection(client *mongo.Client) *mongo.Collection {
	if anwbCollection != nil {
		return anwbCollection
	}
	anwbCollection = client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
	return anwbCollection
}

func Disconnect() {
	if client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client = nil // necessary?
}

func InsertDocument(data anwb.Document) (primitive.ObjectID, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data.Id = primitive.NewObjectID() // necessary?
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func GetDocuments() ([]anwb.Document, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	var entryList []anwb.Document
	for cursor.Next(ctx) {
		var entry anwb.Document
		err := cursor.Decode(&entry)
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}
	return entryList, nil
}

func GetDocumentById(id string) (anwb.Document, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, bson.D{})
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
