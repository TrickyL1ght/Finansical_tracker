package main

import (
	"net/http"
	"strconv"

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
