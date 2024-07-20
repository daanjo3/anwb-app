package handler

import (
	"log/slog"
	"net/http"

	"github.com/daanjo3/anweb-app/api/internal/anwb"
	"github.com/daanjo3/anweb-app/api/internal/db"
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var KEY_DOCUMENT = "document"

// TODO set limit to the amount of resources retrieved by default
// TODO add filtering options
func ListDocuments(c *gin.Context) {
	entries, err := db.GetDocuments(db.DOC_FORMAT_INDEX)
	if err != nil {
		slog.Error("Failed to list documents", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ANWB document index"})
		return
	}
	c.JSON(200, entries)
}

// Retrieves the document that is references in the URL path and sets it in the request context
// to be picked up by other handlers.
func WithDocumentContext(c *gin.Context) {
	var document anwb.Document
	var err error
	id := c.Params.ByName("id")
	if len(id) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID parameter is empty"})
		return
	}
	document, err = service.GetDocument(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Document could not be found"})
			return
		}
		slog.Error("Looking up document failed unexpectedly", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to find document"})
		return
	}
	c.Set(KEY_DOCUMENT, document)
	c.Next()
}

func ReadDocumentById(c *gin.Context) {
	document, exists := c.Get(KEY_DOCUMENT)
	if !exists {
		// Should generally never be reached
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find document"})
		return
	}
	c.JSON(200, document)
}

func UpdateManual(c *gin.Context) {
	document, err := service.AddDocument(false)
	if err != nil {
		slog.Error("Failed to update document manually", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed update ANWB document"})
		return
	}
	c.JSON(201, document)
}
