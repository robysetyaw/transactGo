package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo      repository.AccountRepository
	userRepo         repository.UserRepository
}

func NewTransactionService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: r}
}

func (s *TransactionService) GetTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactions()
}

func (s *TransactionService) GetTransaction(id string) (model.Transaction, error) {
	return s.transactionRepo.GetTransaction(id)
}

func (s *TransactionService) CreateTransaction(tx model.Transaction, username any) error {
	usernameString := username.(string)
	userRn , _ := s.userRepo.FindByUsername(usernameString)
	senderAccount, _ := s.accountRepo.FindByCustomerId(userRn.ID)
	tx.FromAccountNumber = senderAccount.AccountNumber
	return s.transactionRepo.CreateTransaction(tx)
}
