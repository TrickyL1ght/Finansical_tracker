package database

import (
	"awesomeProject/internal/service/transaction"
	"awesomeProject/internal/service/user"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {

		panic(".env file not found")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not init DB: %v", err)
	}

	if err = db.AutoMigrate(&transaction.Transaction{}); err != nil {
		log.Fatalf("could not migrate DB: %v", err)
	}
	if err = db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("could not migrate DB: %v", err)
	}

	return db, nil
}
