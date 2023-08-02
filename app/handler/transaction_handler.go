package handler

import (
	"net/http"
	"transactgo/app/middleware"
	"transactgo/app/model"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService, r *gin.Engine) *TransactionHandler {
	handler := &TransactionHandler{service: s}

	// Set up routes
	r.GET("/transactions",middleware.AuthMiddleware(), handler.GetTransactions)
	r.GET("/transactions/:id", handler.GetTransaction)
	r.POST("/transactions", middleware.AuthMiddleware() , handler.CreateTransaction)

	return handler
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	txs, err := h.service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}
	c.JSON(http.StatusOK, txs)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	tx, err := h.service.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, tx)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var tx model.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	if err := h.service.CreateTransaction(tx,user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	c.JSON(http.StatusCreated, tx)
}
