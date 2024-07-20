package handler

import (
	"fmt"
	"log/slog"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/daanjo3/anweb-app/api/internal/db"
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-gonic/gin"
)

var KEY_DOCUMENT = "document"

// TODO set limit to the amount of resources retrieved by default
// TODO add filtering options
func ListDocuments(c *gin.Context) {
	entries, err := db.GetDocuments(db.DOC_FORMAT_INDEX)
	if err != nil {
		c.Status(500)
		fmt.Fprintf(c.Writer, "Failed to fetch ANWB document index %v", err)
		return
	}
	c.JSON(200, entries)
}

func WithDocumentContext(c *gin.Context) {
	var document anwb.Document
	var err error
	id := c.Params.ByName("id")
	if len(id) == 0 {
		c.AbortWithStatusJSON(404, gin.H{"status": "Document could not be found"})
		return
	}
	document, err = service.GetDocument(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"status": "Document could not be found"})
		slog.Error("Fetching document failed with error\n%+v", err)
		return
	}
	c.Set(KEY_DOCUMENT, document)
	c.Next()
}

func ReadDocumentById(c *gin.Context) {
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		// Should generally never be reached
		c.Status(500)
		fmt.Fprint(c.Writer, "Failed to find document")
	}
	c.JSON(200, document)
}

func UpdateManual(c *gin.Context) {
	document, err := service.AddDocument(false)
	if err != nil {
		c.Status(500)
		fmt.Fprintf(c.Writer, "Failed update ANWB document: %v", err)
	}
	c.JSON(201, document)
}
