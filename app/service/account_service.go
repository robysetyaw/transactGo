package service

import (
	"errors"
	"fmt"
	"transactgo/app/model"
	"transactgo/app/repository"

	"github.com/google/uuid"
)

type AccountService interface {
	FindByAccountNumber(accountNumber string) (*model.Account, error)
	FindAllActive() []model.Account
	Update(account *model.Account) error
	Delete(account *model.Account) error
	Save(account *model.Account) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{
		repo: repo,
	}
}

func (s *accountService) FindAllActive() []model.Account {
	return s.repo.FindAllActive()
}

func (s *accountService) FindByAccountNumber(accountNumber string) (*model.Account, error) {
	return s.repo.FindByAccountNumber(accountNumber)
}

func (s *accountService) Update(account *model.Account) error {
	return s.repo.Update(account)
}

func (s *accountService) Delete(account *model.Account) error {
	return s.repo.Delete(account)
}

func (s *accountService) Save(account *model.Account) error {
	account.AccountNumber = s.generateAccountNumber()
	isExist,_ := s.repo.FindByAccountNumber(account.AccountNumber)
	if isExist != nil {
		return errors.New("account already exist")
	}
	account.ID = uuid.New().String()
	return s.repo.Save(account)
}

func (s *accountService) generateAccountNumber() string {
	count := s.repo.CountAccounts()
	return fmt.Sprintf("00100000%03d", count+1)
}