package main

import (
	"net/http"
	"strconv"
	"time"

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

	if err := NewTransaction(
		transaction.Transaction_type,
		transaction.Category,
		transaction.Description,
		transaction.Amount,
		transaction.Date); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func GetTransactionByIdHandler(c *gin.Context) {

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction type must be int"})
		return
	}

	transaction, err := SelectTransactionById(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func EditTransactionHandler(c *gin.Context) {

	var req Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Date == 0 {
		req.Date = time.Now().Unix()
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction type must be int"})
		return
	}

	transaction, err := SelectTransactionById(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction not found"})
		return
	}
	transaction.Transaction_type = req.Transaction_type
	transaction.Amount = req.Amount
	transaction.Category = req.Category
	transaction.Description = req.Description
	transaction.Date = req.Date

	if err := db.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change data processing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
