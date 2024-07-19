package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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

func Get_Documents() ([]AnwbDoc, error) {
	collection, _, context, cancel := SetupMongoDB()
	defer cancel()
	var entryList []AnwbDoc
	cursor, err := collection.Find(context, bson.D{})
	defer cursor.Close(context)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context) {
		var entry AnwbDoc
		err := cursor.Decode(&entry)
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}
	return entryList, nil
}

func Get_Document_ById(id string) (AnwbDoc, error) {
	collection, _, context, cancel := SetupMongoDB()
	defer cancel()
	result := collection.FindOne(context, bson.D{})
	if result.Err() != nil {
		return AnwbDoc{}, result.Err()
	}
	var entry AnwbDoc
	err := result.Decode(&entry)
	if err != nil {
		return AnwbDoc{}, err
	}
	return entry, nil
}
