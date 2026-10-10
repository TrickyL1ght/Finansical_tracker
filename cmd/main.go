package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/service/transaction"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("Database could not init")
	}

	repository := transaction.NewTransactionRepository(db)
	service := transaction.NewTransactionService(repository)
	handler := handlers.NewTransactionHandler(service)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://finansicaltracker.framer.website", "http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "User-Agent", "Sec-Ch-Ua-Platform", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "origin", "host", "connection", "accept-language", "accept-encoding", "accept"},
	}))

	router.GET("/api/transaction", handler.GetAllTransaction)
	router.POST("/api/transaction", handler.CreateTransaction)
	router.GET("/api/transaction/:id", handler.GetTransactionById)
	router.PUT("/api/transaction/:id", handler.UpdateTransaction)
	router.DELETE("/api/transaction/:id", handler.DeleteTransaction)

	router.Run(":8080")
}
