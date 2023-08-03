package handler

import (
	"net/http"
	"strconv"
	"transactgo/app/model"
	"transactgo/app/model/response"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service service.MerchantService
}

func NewMerchantHandler(s service.MerchantService, r *gin.Engine) *MerchantHandler {
	handler := &MerchantHandler{service: s}

	r.GET("/merchants/:id", handler.GetMerchant)
	r.POST("/merchants", handler.CreateMerchant)
	r.PUT("/merchants/:id", handler.UpdateMerchant)
	r.DELETE("/merchants/:id", handler.DeleteMerchant)
	r.GET("/merchants", handler.GetAllMerchants)

	return handler
}

func (h *MerchantHandler) GetMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", "Invalid merchant ID", nil, err.Error()))
		return
	}

	merchant, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Error occurred while fetching the merchant", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get merchant", merchant, " "))
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", err.Error(), nil, ""))
		return
	}
	user, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to get user", nil, ""))
		return
	}
	username := user.(string)
	if err := h.service.CreateMerchant(&merchant, username ); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Error occurred while creating the merchant", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully added merchant", merchant, " "))
}

func (h *MerchantHandler) UpdateMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", "Invalid merchant ID", nil, err.Error()))
		return
	}

	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", err.Error(), nil, ""))
		return
	}

	merchant.ID = id
	if err := h.service.UpdateMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Error occurred while updating the merchant", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully updated merchant", merchant, " "))
}

func (h *MerchantHandler) DeleteMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", "Invalid merchant ID", nil, err.Error()))
		return
	}

	if err := h.service.DeleteMerchant(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Error occurred while deleting the merchant", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "Merchant deleted successfully", nil, " "))
}

func (h *MerchantHandler) GetAllMerchants(c *gin.Context) {
	merchants, err := h.service.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Error occurred while fetching all merchants", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get all merchants", merchants, " "))
}
