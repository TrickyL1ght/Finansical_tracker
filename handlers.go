package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactionHandler(c *gin.Context) {
	allTransaction, err := SelectAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, allTransaction)
}

func CreateTransactionHandler(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := NewTransaction(transaction.Transaction_type,
		transaction.Category,
		transaction.Description,
		transaction.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

//
//func EditTransactionHandler(c *gin.Context) {
//	path := strings.TrimPrefix(r.URL.Path, "/api/transactions/")
//	transactionId := path
//	switch r.Method {
//	case http.MethodDelete:
//		if err := db.Delete(&Transaction{}, "id = ?", transactionId).Error; err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
//			json.NewEncoder(w).Encode(err)
//		}
//		w.WriteHeader(http.StatusNoContent)
//	}
//}
