package handler

import (
	"net/http"
	"transactgo/app/middleware"
	"transactgo/app/model"
	"transactgo/app/model/response"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService, r *gin.Engine) *TransactionHandler {
	handler := &TransactionHandler{service: s}

	// Set up routes
	r.GET("/transactions", middleware.AuthMiddleware(), handler.GetTransactions)
	r.GET("/transactions/:id", handler.GetTransaction)
	r.POST("/transactions", middleware.AuthMiddleware(), handler.CreateTransaction)

	return handler
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	txs, err := h.service.GetTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to get transactions", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get transactions", txs, ""))
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	tx, err := h.service.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", "Transaction not found", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get transaction", tx, ""))
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var tx model.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", err.Error(), nil, ""))
		return
	}

	user, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to get user", nil, ""))
		return
	}

	data, err := h.service.CreateTransaction(tx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to create transaction", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.NewResponse(http.StatusCreated, "OK","Successfully created transaction", data, ""))
}
