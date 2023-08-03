package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type MerchantService interface {
	GetByID(id int) (*model.Merchant, error)
	CreateMerchant(merchant *model.Merchant, username string) error
	UpdateMerchant(merchant *model.Merchant) error
	DeleteMerchant(id int) error
	GetAllMerchants() ([]model.Merchant, error)
}

type merchantService struct {
	repo     repository.MerchantRepository
	userRepo repository.UserRepository
}

func NewMerchantService(repo repository.MerchantRepository, userRepo repository.UserRepository) MerchantService {
	return &merchantService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *merchantService) GetByID(id int) (*model.Merchant, error) {
	return s.repo.GetByID(id)
}

func (s *merchantService) CreateMerchant(merchant *model.Merchant, username string) error {
	user, _ := s.userRepo.FindByUsername(username)
	merchant.UserID = user.ID
	return s.repo.Save(merchant)
}

func (s *merchantService) UpdateMerchant(merchant *model.Merchant) error {
	return s.repo.Update(merchant)
}

func (s *merchantService) DeleteMerchant(id int) error {
	return s.repo.Delete(id)
}

func (s *merchantService) GetAllMerchants() ([]model.Merchant, error) {
	return s.repo.FindAll()
}
