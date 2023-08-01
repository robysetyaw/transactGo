package handler

import (
	"net/http"
	// "transactgo/app/model"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService, r *gin.Engine) *UserHandler {
	handler := &UserHandler{service: s}

	// Set up routes
	r.GET("/users/:username", handler.GetUserByUsername)
	r.PUT("/users/:username", handler.UpdateUser)
	r.DELETE("/users/:username", handler.DeleteUser)

	return handler
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user := h.service.GetUserByUsername(username)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    username := c.Param("username")
    existingUser := h.service.GetUserByUsername(username)
    if existingUser == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Define a new struct to hold the update request
    type UpdateRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    var update UpdateRequest
    if err := c.ShouldBindJSON(&update); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the user's username and password
    existingUser.Username = update.Username
    existingUser.Password = update.Password

    if err := h.service.UpdateUser(existingUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }
    c.JSON(http.StatusOK, existingUser)
}


func (h *UserHandler) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	user := h.service.GetUserByUsername(username)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := h.service.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "User deleted"})
}