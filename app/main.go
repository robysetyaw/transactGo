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
	userRepo, _ := repository.NewUserRepository()
	accountRepo , _ := repository.NewAccountRepository()
	merchantRepo, _ := repository.NewMerchantRepository()

	// Create UserService
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo)
	merchantService := service.NewMerchantService(merchantRepo)

	// Create UserHandler
	handler.NewUserHandler(userService, r)
	handler.NewAccountHandler(accountService, r)
	handler.NewMerchantHandler(merchantService, r)
	
	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
