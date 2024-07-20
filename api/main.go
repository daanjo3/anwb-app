package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/daanjo3/anweb-app/api/internal/db"
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
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

	c := cron.New()
	// TODO fix cronjob, is currently every 5 seconds..
	c.AddFunc("*/5 * * * * *", func() {
		_, err := service.UpdateLocal()
		if err != nil {
			slog.Error("Failed to perform automatic update", "error", err)
		}
		slog.Info("Automatically updated document")
	})
	c.Start()

	r := gin.Default()
	r.GET("/documents", service.GetDocuments)
	r.GET("/documents/:id", service.GetDocumentById)
	r.POST("/update", service.UpdateManual)
	r.Run()
}
