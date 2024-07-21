package db

import (
	"context"
	"log"
	"os"
	"time"

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
	client = nil // TODO check if required
}
