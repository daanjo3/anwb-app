package service

import (
	"fmt"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/daanjo3/anweb-app/api/internal/db"
	"github.com/gin-gonic/gin"
)

// TODO set limit to the amount of resources retrieved by default
// TODO add filtering options
func GetDocuments(c *gin.Context) {
	entries, err := db.Get_Documents()
	if err != nil {
		fmt.Fprintf(c.Writer, "Failed to fetch ANWB document index %v", err)
		c.Status(500)
	}
	c.JSON(200, entries)
}

func GetDocumentById(c *gin.Context) {
	id := c.Params.ByName("id")
	if len(id) == 0 {
		c.Status(400)
		fmt.Fprintf(c.Writer, "Could not find document with ID %v", id)
	}
	document, err := db.Get_Document_ById(id)
	if err != nil {
		c.Status(404)
		fmt.Fprintf(c.Writer, "Could not find document with ID %v", id)
	}
	c.JSON(200, document)
}

func UpdateManual(c *gin.Context) {
	document, err := update()
	if err != nil {
		fmt.Fprintf(c.Writer, "Failed update ANWB document: %v", err)
		c.Status(500)
	}
	c.JSON(201, document)
}

// TODO specify the errors
func update() (anwb.Document, error) {
	data, err := anwb.Get()
	if err != nil {
		return anwb.Document{}, err
	}
	id, err := db.InsertDocument(data)
	if err != nil {
		return anwb.Document{}, err
	}
	fmt.Printf("Inserted new ANWB document with id %v", id)
	return data, nil
}
