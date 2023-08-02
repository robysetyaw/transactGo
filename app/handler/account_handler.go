package handler

import (
	"net/http"
	"transactgo/app/model"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService, r *gin.Engine) *AccountHandler {
	h := &AccountHandler{
		accountService: accountService,
	}
	r.GET("/accounts/:accountNumber", h.GetAccount)
	r.POST("/accounts", h.CreateAccount)
	r.PUT("/accounts/:accountNumber", h.UpdateAccount)
	r.DELETE("/accounts/:accountNumber", h.DeleteAccount)
	return h
}



func (h *AccountHandler) GetAccount(c *gin.Context) {
	accountNumber := c.Param("accountNumber")
	account, err := h.accountService.FindByAccountNumber(accountNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var newAccount model.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.accountService.Save(&newAccount)
	c.JSON(http.StatusCreated, newAccount)
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	// accountNumber := c.Param("accountNumber")
	var updatedAccount model.Account
	if err := c.ShouldBindJSON(&updatedAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.Update(&updatedAccount); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedAccount)
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	accountNumber := c.Param("accountNumber")
	account, err := h.accountService.FindByAccountNumber(accountNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := h.accountService.Delete(account); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
