package main

import (
	"transactgo/app/handler"
	"transactgo/app/repository"
	"transactgo/app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create UserRepository
	userRepo, err := repository.NewUserRepository()
	if err != nil {
		panic(err)
	}

	// Create UserService
	userService := service.NewUserService(userRepo)

	// Create UserHandler
	userHandler := handler.NewUserHandler(userService)

	// Set up Gin server
	r := gin.Default()

	// Set up routes
	r.GET("/users/:username", userHandler.GetUserByUsername)
	// Add other routes (e.g., POST, PUT, DELETE) as needed

	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
