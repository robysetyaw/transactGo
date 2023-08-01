package handler

import (
	"net/http"
	"transactgo/app/model"
	"transactgo/app/model/response"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	r.POST("/users",handler.AddUser)
	return handler
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user,_ := h.service.GetUserByUsername(username)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK"," ", user, " "))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    username := c.Param("username")

    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateUser(username, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }
    c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","User with username " + username + " has been updated" ,user, " "))
}


func (h *UserHandler) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	
	if err := h.service.DeleteUser(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "User with username " + username + " has been deleted", " "," "))
}

func (h *UserHandler) AddUser(c *gin.Context) {

    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
	user.Password = string(hashedPassword)
    if err := h.service.AddUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
        return
    }
    c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully added user", user, " "))
}