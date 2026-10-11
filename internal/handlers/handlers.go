package handlers

import (
	"awesomeProject/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://finansicaltracker.framer.website", "http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "User-Agent", "Sec-Ch-Ua-Platform", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "origin", "host", "connection", "accept-language", "accept-encoding", "accept"},
	}))

	transactionApi := router.Group("/api/transaction", h.UserIdentity)
	{
		transactionApi.GET("/list", h.GetAllTransaction)
		transactionApi.POST("", h.CreateTransaction)
		transactionApi.GET("/:id", h.GetTransactionById)
		transactionApi.PUT("/:id", h.UpdateTransaction)
		transactionApi.DELETE("/:id", h.DeleteTransaction)
	}

	userApi := router.Group("/api/user")
	{
		userApi.POST("", h.CreateUser)
		userApi.POST("auth", h.UserAuth)
	}
	return router
}
