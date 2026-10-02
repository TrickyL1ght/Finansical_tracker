package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=127.0.0.1 user=postgres password=HomePass dbname=ftdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not init DB: %v", err)
	}

	if err = db.AutoMigrate(&Transaction{}); err != nil {
		log.Fatalf("could not migrate DB: %v", err)
	}
}

func main() {
	initDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://finansicaltracker.framer.website"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "User-Agent", "Sec-Ch-Ua-Platform", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "origin", "host", "connection", "accept-language", "accept-encoding", "accept"},
	}))

	router.GET("/api/transaction", GetTransactionHandler)
	router.POST("/api/transaction", CreateTransactionHandler)
	router.GET("/api/transaction/:id", GetTransactionByIdHandler)
	router.Run(":8080")
}
