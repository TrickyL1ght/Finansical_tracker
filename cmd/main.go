package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic("Database could not init")
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)

	router := handler.InitRoutes()

	router.RunTLS(":8080", "localhost+1.pem", "localhost+1-key.pem")
}
