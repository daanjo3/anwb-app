package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Update() error {
	data, err := GetAnwbData()
	if err != nil {
		return err
	}
	var bdoc interface{}
	err = bson.UnmarshalExtJSON(data, true, &bdoc)
	collection, client, context, cancel := SetupMongoDB()
	defer CloseConnection(client, context, cancel)

	bdoc.id = primitive.NewObjectID()
	collection.InsertOne(context, bdoc)

}
