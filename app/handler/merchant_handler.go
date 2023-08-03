package handler

import (
	"net/http"
	"strconv"
	"transactgo/app/model"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	merchant, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching the merchant"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, exists := c.Get("username")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	username := user.(string)
	if err := h.service.CreateMerchant(&merchant, username ); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while creating the merchant"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *MerchantHandler) UpdateMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant.ID = id
	if err := h.service.UpdateMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating the merchant"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

func (h *MerchantHandler) DeleteMerchant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	if err := h.service.DeleteMerchant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the merchant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Merchant deleted successfully"})
}

func (h *MerchantHandler) GetAllMerchants(c *gin.Context) {
	merchants, err := h.service.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching all merchants"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}
