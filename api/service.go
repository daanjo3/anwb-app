package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDocumentById(c *gin.Context) {
	id := c.Params.ByName("id")
	if len(id) == 0 {
		c.Status(400)
		fmt.Fprintf(c.Writer, "Could not find document with ID %v", id)
	}
	document, err := Get_Document_ById(id)
	if err != nil {
		c.Status(404)
		fmt.Fprintf(c.Writer, "Could not find document with ID %v", id)
	}
	c.JSON(200, document)
}

func GetDocuments(c *gin.Context) {
	entries, err := Get_Documents()
	if err != nil {
		fmt.Fprintf(c.Writer, "Failed to fetch ANWB document index %v", err)
		c.Status(500)
	}
	c.JSON(200, entries)
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
