package main

import (
	"fmt"

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
	r.GET("/documents", GetDocuments)
	r.GET("/documents/:id", GetDocumentById)
	r.POST("/update", func(c *gin.Context) {
		document, err := Update()
		if err != nil {
			fmt.Printf("Failed to update ANWB data %v", err)
			c.Status(500)
			return
		}
		c.JSON(200, document)
	})
	r.Run()
}
