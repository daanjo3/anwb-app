package db

import (
	"context"
	"time"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ExistsDocument(data *anwb.Document) (bool, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, bson.D{{Key: "_uploaded_at", Value: data.UploadedAt}})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
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

type DocumentFormat int

const (
	DOC_FORMAT_FULL DocumentFormat = iota
	DOC_FORMAT_INDEX
)

func GetDocuments(format DocumentFormat) ([]anwb.Document, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find()
	if format == DOC_FORMAT_INDEX {
		opts.SetProjection(bson.D{{"_id", 1}, {"_uploaded_at", 1}}).SetSort(bson.D{{"_uploaded_at", -1}})
	}
	cursor, err := collection.Find(ctx, bson.D{}, opts)
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
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return anwb.Document{}, err
	}
	return GetDocument(bson.D{{"_id", objectId}}, options.FindOne())
}

func GetDocumentLatest() (anwb.Document, error) {
	return GetDocument(bson.D{{}}, options.FindOne().SetSort(bson.D{{"_uploaded_at", -1}}))
}

func GetDocument(filter bson.D, opts *options.FindOneOptions) (anwb.Document, error) {
	collection := GetAnwbCollection(GetClient())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, filter, opts)
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
