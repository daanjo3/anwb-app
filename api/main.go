package main

import (
	"github.com/daanjo3/anweb-app/api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// c := cron.New()
	// c.AddFunc("*/5 * * * * *", func() {
	// 	_, err := Update()
	// 	if err != nil {
	// 		fmt.Printf("Failed to update ANWB data %v", err)
	// 	}
	// })
	// c.Start()

	r := gin.Default()
	r.GET("/documents", service.GetDocuments)
	r.GET("/documents/:id", service.GetDocumentById)
	r.POST("/update", service.UpdateManual)
	r.Run()
}
