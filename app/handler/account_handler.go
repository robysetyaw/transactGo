package handler

import (
	"net/http"
	"transactgo/app/middleware"
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
	r.GET("/accounts", h.GetActiveAccounts)
	r.POST("/accounts",middleware.AuthMiddleware() , h.CreateAccount)
	r.PUT("/accounts/:accountNumber", h.UpdateAccount)
	r.DELETE("/accounts/:accountNumber", h.DeleteAccount)
	return h
}

func (h *AccountHandler) GetActiveAccounts(c *gin.Context) {
	activeAccounts := h.accountService.FindAllActive()
	c.JSON(http.StatusOK, activeAccounts)
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
	user, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	userString := user.(string)
	h.accountService.Save(&newAccount,userString)
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
	if err := h.accountService.DeActive(account); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
