package service

import (
	"fmt"
	"time"
	"transactgo/app/model"
	"transactgo/app/repository"

	"github.com/google/uuid"
)

type TransactionService interface {
	GetTransactions() ([]model.Transaction, error)
	GetTransaction(id string) (model.Transaction, error)
	CreateTransaction(tx model.Transaction, username any) (model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
	userRepo        repository.UserRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, accountRepo repository.AccountRepository, userRepo repository.UserRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		userRepo:        userRepo,
	}
}

func (s *transactionService) GetTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactions()
}

func (s *transactionService) GetTransaction(id string) (model.Transaction, error) {
	return s.transactionRepo.GetTransaction(id)
}

func (s *transactionService) CreateTransaction(tx model.Transaction, username any) (model.Transaction ,error) {
	usernameString := username.(string)
	userRn , _ := s.userRepo.FindByUsername(usernameString)
	senderAccount, _ := s.accountRepo.FindByCustomerId(userRn.ID)
	receiverAccount, _ := s.accountRepo.FindByAccountNumber(tx.ToAccountNumber)
	if receiverAccount.IsMerchant {
		tx.TxType = "payment"
	}
	if !receiverAccount.IsMerchant {
		tx.TxType = "transfer"
	}
	tx.FromAccountNumber = senderAccount.AccountNumber
	tx.ID = uuid.New().String()
	tx.Timestamp = time.Now()

	txNumber, _ := s.generateTxNumber(tx.TxType)
	tx.TxNumber = txNumber
	s.transactionRepo.CreateTransaction(tx)
	return tx,nil
}

func (s *transactionService) generateTxNumber(txType string) (string, error) {

	txs, err := s.transactionRepo.GetTransactions()
	if err != nil {
		return "", err
	}
	count := 0
	today := time.Now().Format("20060102") // Format tanggal ke bentuk YYYYMMDD
	for _, tx := range txs {
		if tx.Timestamp.Format("20060102") == today {
			count++
		}
	}
	prefix := ""
	if txType == "payment" {
		prefix = "PY"
	} else if txType == "transfer" {
		prefix = "TX"
	}
	return fmt.Sprintf("%s%s%04d", prefix, today, count+1), nil
}
