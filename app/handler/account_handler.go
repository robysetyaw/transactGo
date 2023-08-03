package handler

import (
	"net/http"
	"transactgo/app/middleware"
	"transactgo/app/model"
	"transactgo/app/model/response"
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
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get active accounts", activeAccounts, " "))
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	accountNumber := c.Param("accountNumber")
	account, err := h.accountService.FindByAccountNumber(accountNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", err.Error(), nil, ""))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get account", account, " "))
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var newAccount model.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", err.Error(), nil, ""))
		return
	}
	user, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to get user", nil, ""))
		return
	}
	userString := user.(string)
	h.accountService.Save(&newAccount,userString)
	c.JSON(http.StatusCreated, response.NewResponse(http.StatusCreated, "OK","Successfully created account", newAccount, " "))
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	var updatedAccount model.Account
	if err := c.ShouldBindJSON(&updatedAccount); err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", err.Error(), nil, ""))
		return
	}
	if err := h.accountService.Update(&updatedAccount); err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", err.Error(), nil, ""))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully updated account", updatedAccount, " "))
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	accountNumber := c.Param("accountNumber")
	account, err := h.accountService.FindByAccountNumber(accountNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", err.Error(), nil, ""))
		return
	}
	if err := h.accountService.DeActive(account); err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", err.Error(), nil, ""))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Account deleted successfully", nil, " "))
}
