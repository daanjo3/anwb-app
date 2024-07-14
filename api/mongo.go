package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() (*mongo.Client, error, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))
	if err != nil {
		return nil, err, nil
	}

	cancelClient := func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return client, err, cancelClient
}

func Update() error {
	data, err := GetAnwbData()
	if err != nil {
		return err
	}
	var bdoc interface{}
	err = bson.UnmarshalExtJSON(data, true, &bdoc)

	client, error, cancel := GetClient()
	if error != nil {
		return err
	}
	defer cancel()
	collection := client.Database("testing").Collection("numbers")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID
}
