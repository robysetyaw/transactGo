package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transactgo/app/model"
)

type AccountRepository interface {
	FindByAccountNumber(accountNumber string) (*model.Account, error)
	Update(account *model.Account) error
	Delete(account *model.Account) error
	Save(account *model.Account) error
	CountAccounts() int
	FindAllActive() []model.Account
}

type accountRepository struct {
	accounts []model.Account
}

func NewAccountRepository() (AccountRepository, error) {
	repo := &accountRepository{}

	// Open the JSON file
	file, err := os.Open("data/accounts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the file into the accounts slice
	err = json.NewDecoder(file).Decode(&repo.accounts)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *accountRepository) FindByAccountNumber(accountNumber string) (*model.Account, error) {
	for _, account := range r.accounts {
		if account.AccountNumber == accountNumber {
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}

func (r *accountRepository) Update(account *model.Account) error {
	// Update the account in the slice
	for i, a := range r.accounts {
		if a.AccountNumber == account.AccountNumber {
			r.accounts[i] = *account
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/accounts.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the accounts slice back into the file
	err = json.NewEncoder(file).Encode(r.accounts)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) Delete(account *model.Account) error {
	// Remove the account from the slice
	for i, a := range r.accounts {
		if a.AccountNumber == account.AccountNumber {
			r.accounts = append(r.accounts[:i], r.accounts[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/accounts.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the accounts slice back into the file
	err = json.NewEncoder(file).Encode(r.accounts)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) Save(account *model.Account) error {
	// Check if the account exists in the slice
	found := false
	for i, a := range r.accounts {
		if a.AccountNumber == account.AccountNumber {
			r.accounts[i] = *account
			found = true
			break
		}
	}

	// If the account does not exist, add it to the slice
	if !found {
		r.accounts = append(r.accounts, *account)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/accounts.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the accounts slice back into the file
	err = json.NewEncoder(file).Encode(r.accounts)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepository) CountAccounts() int {
	return len(r.accounts)
}

func (r *accountRepository) FindAllActive() []model.Account {
	var activeAccounts []model.Account
	for _, account := range r.accounts {
		if account.IsActive {
			activeAccounts = append(activeAccounts, account)
		}
	}
	return activeAccounts
}