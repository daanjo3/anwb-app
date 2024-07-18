package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListDocuments() ([]IndexEntry, error) {
	collection, _, context, cancel := SetupMongoDB()
	defer cancel()
	var entryList []IndexEntry
	cursor, err := collection.Find(context, bson.D{})
	defer cursor.Close(context)
	if err != nil {
		return nil, err
	}
	//  TODO make iterating the entries possible
	for cursor.Next(context) {
		var entry IndexEntry
		err := cursor.Decode(&entry)
		fmt.Printf("Entry: %+v", entry)
		if err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}
	return entryList, nil
}

func Update() (AnwbDoc, error) {
	data, err := GetAnwbDocument()
	if err != nil {
		return AnwbDoc{}, err
	}
	collection, client, context, cancel := SetupMongoDB()
	defer CloseConnection(client, context, cancel)

	data.Id = primitive.NewObjectID() // necessary?
	result, err := collection.InsertOne(context, data)
	if err != nil {
		return AnwbDoc{}, err
	}
	fmt.Printf("Inserted document %v", result.InsertedID)
	return data, nil
}
