package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) GetTransactions() ([]model.Transaction, error) {
	return s.repo.GetTransactions()
}

func (s *TransactionService) GetTransaction(id string) (model.Transaction, error) {
	return s.repo.GetTransaction(id)
}

func (s *TransactionService) CreateTransaction(tx model.Transaction) error {
	return s.repo.CreateTransaction(tx)
}
