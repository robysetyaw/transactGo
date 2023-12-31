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
	DeActive(account *model.Account) error
	Save(account *model.Account, username string) error
}

type accountService struct {
	repo repository.AccountRepository
	userRepo repository.UserRepository
}

func NewAccountService(repo repository.AccountRepository, userRepo repository.UserRepository) AccountService {
	return &accountService{
		repo: repo,
		userRepo: userRepo,
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

func (s *accountService) DeActive(account *model.Account) error {
	account.IsActive = false
	return s.repo.Update(account)
}

func (s *accountService) Save(account *model.Account, username string) error {
	user, _ := s.userRepo.FindByUsername(username)
	account.UserID = user.ID
	account.AccountNumber = s.generateAccountNumber()
	isExist,_ := s.repo.FindByAccountNumber(account.AccountNumber)
	if isExist != nil {
		return errors.New("account already exist")
	}
	account.ID = uuid.New().String()
	account.IsActive = true
	return s.repo.Save(account)
}

func (s *accountService) generateAccountNumber() string {
	count := s.repo.CountAccounts()
	return fmt.Sprintf("00100000%03d", count+1)
}