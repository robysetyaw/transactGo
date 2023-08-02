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
	accountRepo , _ := repository.NewAccountRepository()

	// Create UserService
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo)

	// Create UserHandler
	handler.NewUserHandler(userService, r)
	handler.NewAccountHandler(accountService, r)
	
	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
