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
	accountRepo, _ := repository.NewAccountRepository()
	merchantRepo, _ := repository.NewMerchantRepository()
	transactionRepo, _ := repository.NewTransactionRepository()

	// Create UserService
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo,userRepo)
	merchantService := service.NewMerchantService(merchantRepo)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo, userRepo)

	// Create UserHandler
	handler.NewUserHandler(userService, r)
	handler.NewAccountHandler(accountService, r)
	handler.NewMerchantHandler(merchantService, r)
	handler.NewTransactionHandler(transactionService, r)
	
	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
