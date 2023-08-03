package handler

import (
	"net/http"
	"time"
	"transactgo/app/middleware"
	"transactgo/app/model"
	"transactgo/app/model/response"
	"transactgo/app/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

type SafeUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewUserHandler(s service.UserService, r *gin.Engine) *UserHandler {
	handler := &UserHandler{service: s}

	r.GET("/users/:username", middleware.AuthMiddleware(), handler.GetUserByUsername)
	r.PUT("/users/:username", middleware.AuthMiddleware(), handler.UpdateUser)
	r.DELETE("/users/:username", middleware.AuthMiddleware(), handler.DeleteUser)
	r.POST("/users", handler.AddUser)
	r.POST("/login", handler.Login)
	return handler
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewResponse(http.StatusNotFound, "Failed", "User not found", nil, ""))
		return
	}

	safeUser := SafeUser{
		ID:       user.ID,
		Username: user.Username,
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "Successfully get user", safeUser, ""))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    username := c.Param("username")

    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Failed", err.Error(), nil, ""))
        return
    }

    if err := h.service.UpdateUser(username, &user); err != nil {
        c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to update user", nil, ""))
        return
    }
    c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "User with username " + username + " has been updated", nil, ""))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	
	if err := h.service.DeleteUser(username); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to delete user", nil, ""))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "User with username " + username + " has been deleted", nil, ""))
}

func (h *UserHandler) AddUser(c *gin.Context) {
    var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Failed", err.Error(), nil, ""))
        return
    }

	errors := make(map[string]string)

	if len(user.Username) < 5 {
		errors["username"] = "username must be 5 characters or more"
	}

	if len(user.Password) < 6 {
		errors["password"] = "password must be 6 characters or more"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", errors, nil, ""))
		return
	}

	safeUser := SafeUser{
		ID:       user.ID,
		Username: user.Username,
	}

    if err := h.service.AddUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Failed to add user", nil, ""))
        return
    }
    c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK", "Successfully added user", safeUser , ""))
}

func (h *UserHandler) Login(c *gin.Context) {
    var loginReq model.User
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Failed", err.Error(), nil, ""))
        return
    }

    user, err := h.service.Authenticate(loginReq.Username, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, response.NewResponse(http.StatusUnauthorized, "Failed", "Invalid username or password", nil, ""))
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
		"exp":      time.Now().Add(time.Hour * 10).Unix(),
    })

    tokenString, err := token.SignedString([]byte("secret.puppey")) 
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed", "Could not generate token", nil, ""))
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
