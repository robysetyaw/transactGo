package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transactgo/app/model"
)

type TransactionRepository interface {
	GetTransactions() ([]model.Transaction, error)
	GetTransaction(id string) (model.Transaction, error)
	CreateTransaction(tx model.Transaction) error
}

type transactionRepository struct {
	transactions []model.Transaction
}

func NewTransactionRepository() TransactionRepository {
	repo := &transactionRepository{}

	// Open the JSON file
	file, err := os.Open("data/transactions.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the file into the transactions slice
	err = json.NewDecoder(file).Decode(&repo.transactions)
	if err != nil {
		return nil
	}

	return repo
}

func (r *transactionRepository) GetTransactions() ([]model.Transaction, error) {
	return r.transactions, nil
}

func (r *transactionRepository) GetTransaction(id string) (model.Transaction, error) {
	for _, tx := range r.transactions {
		if tx.ID == id {
			return tx, nil
		}
	}
	return model.Transaction{}, errors.New("Transaction not found")
}

func (r *transactionRepository) CreateTransaction(tx model.Transaction) error {
	r.transactions = append(r.transactions, tx)

	// Open the JSON file
	file, err := os.OpenFile("data/transactions.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the transactions slice back into the file
	err = json.NewEncoder(file).Encode(r.transactions)
	if err != nil {
		return err
	}

	return nil
}
