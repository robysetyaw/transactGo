package handler

import (
	"net/http"
	"transactgo/app/middleware"
	"transactgo/app/model"
	"transactgo/app/model/response"
	"transactgo/app/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService, r *gin.Engine) *UserHandler {
	handler := &UserHandler{service: s}

	// Set up routes
	r.GET("/users/:username", middleware.AuthMiddleware(),handler.GetUserByUsername)
	r.PUT("/users/:username", middleware.AuthMiddleware(),handler.UpdateUser)
	r.DELETE("/users/:username",middleware.AuthMiddleware(), handler.DeleteUser)
	r.POST("/users",handler.AddUser)
	r.POST("/login", handler.Login)
	return handler
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user,_ := h.service.GetUserByUsername(username)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully get user", user, " "))
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

	errors := make(map[string]string)

	// validasi username
	if len(user.Username) < 5 {
		errors["username"] = "username must be 5 characters or more"
	}

	// validasi password
	if len(user.Password) < 6 {
		errors["password"] = "password must be 6 characters or more"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, response.NewResponse(http.StatusBadRequest, "Invalid input", errors, nil, ""))
		return
	}

    if err := h.service.AddUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, response.NewResponse(http.StatusInternalServerError, "Failed to add user", nil, nil, err.Error()))
        return
    }
    c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "OK","Successfully added user", user, " "))
}

func (h *UserHandler) Login(c *gin.Context) {
    var loginReq model.User
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.Authenticate(loginReq.Username, loginReq.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
    })

    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString([]byte("secret.puppey"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
