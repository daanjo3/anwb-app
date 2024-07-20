package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/daanjo3/anweb-app/api/internal/db"
	"github.com/daanjo3/anweb-app/api/internal/handler"
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// assert DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := db.GetClient()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()
	slog.Info("Connected to MongoDB")

	// cronjob to run update every 5 minutes
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() {
		_, err := service.AddDocument(true)
		if err != nil {
			slog.Error("Failed to perform automatic update", "error", err)
		}
		slog.Info("Automatically updated document")
	})
	c.Start()

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/documents", handler.ListDocuments)
	r.GET("/documents/:id", handler.WithDocumentContext, func(c *gin.Context) {
		document, exists := c.Get(handler.KEY_DOCUMENT)
		if !exists {
			// Should generally never be reached
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find document"})
			return
		}
		c.JSON(200, document)
	})
	r.GET("/documents/:id/events/jams", handler.WithDocumentContext, handler.ListJams)
	r.GET("/documents/:id/events/roadworks", handler.WithDocumentContext, handler.ListRoadWorks)
	r.GET("/documents/:id/events/radars", handler.WithDocumentContext, handler.ListRadars)
	r.POST("/update", handler.UpdateManual)
	r.Run()
}
