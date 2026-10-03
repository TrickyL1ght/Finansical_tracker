package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SSL_MODE"),
	)

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

	if err := godotenv.Load(); err != nil {
		panic(".env file not found")
	}

	initDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://finansicaltracker.framer.website", "http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "User-Agent", "Sec-Ch-Ua-Platform", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "origin", "host", "connection", "accept-language", "accept-encoding", "accept"},
	}))

	router.GET("/api/transaction", GetTransactionHandler)
	router.POST("/api/transaction", CreateTransactionHandler)
	router.GET("/api/transaction/:id", GetTransactionByIdHandler)
	router.PUT("/api/transaction/:id", EditTransactionHandler)

	router.Run(":8080")
}
