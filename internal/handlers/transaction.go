package handlers

import (
	"awesomeProject/internal/service/transaction"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service transaction.TransactionService
}

func NewTransactionHandler(service transaction.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetAllTransaction(c *gin.Context) {
	allTransaction, err := h.service.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, allTransaction)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var request transaction.Transaction
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.service.CreateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, transaction)
}

func (h *TransactionHandler) GetTransactionById(c *gin.Context) {

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction type must be int"})
		return
	}

	transaction, err := h.service.GetTransactionById(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {

	var request transaction.Transaction
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction type must be int"})
		return
	}

	request.Id = transactionId

	transaction, err := h.service.UpdateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction type must be int"})
		return
	}

	if err := h.service.DeleteTransaction(transactionId); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
