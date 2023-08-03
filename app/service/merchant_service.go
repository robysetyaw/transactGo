package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type MerchantService struct {
	repo repository.MerchantRepository
	userRepo repository.UserRepository
}

func NewMerchantService(repo repository.MerchantRepository, userRepo repository.UserRepository) *MerchantService {
	return &MerchantService{
		repo: repo,
		userRepo: userRepo,
	}
}

func (s *MerchantService) GetByID(id int) (*model.Merchant, error) {
	return s.repo.GetByID(id)
}

func (s *MerchantService) CreateMerchant(merchant *model.Merchant, username string) error {
	user, _ := s.userRepo.FindByUsername(username)
	merchant.UserID = user.ID
	return s.repo.Save(merchant)
}

func (s *MerchantService) UpdateMerchant(merchant *model.Merchant) error {
	return s.repo.Update(merchant)
}

func (s *MerchantService) DeleteMerchant(id int) error {
	return s.repo.Delete(id)
}

func (s *MerchantService) GetAllMerchants() ([]model.Merchant, error) {
	return s.repo.FindAll()
}
