package main

import (
	"log"
	"net/http"

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

	http.HandleFunc("/api/transaction", TransactionHandler)
	http.HandleFunc("/api/transaction/:tranasationId", EditTransactionHandler)

	http.ListenAndServe(":8080", nil)
}
