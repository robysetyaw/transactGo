package main

import (
	"transactgo/app/handler"
	"transactgo/app/repository"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set up Gin server
	r := gin.Default()
	// Create UserRepository
	userRepo, err := repository.NewUserRepository()
	if err != nil {
		panic(err)
	}

	// Create UserService
	userService := service.NewUserService(userRepo)

	// Create UserHandler
	handler.NewUserHandler(userService, r)
	
	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
