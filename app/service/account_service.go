package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type AccountService interface {
	FindByAccountNumber(accountNumber string) (*model.Account, error)
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
	return s.repo.Save(account)
}
