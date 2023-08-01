package service

import (
	"errors"
	"transactgo/app/model"
	"transactgo/app/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByUsername(username string) *model.User {
	return s.repo.FindByUsername(username)
}

func (s *UserService) UpdateUser(username string, user *model.User) error {
	err, isExist := s.repo.FindByUsername(username)
	if isExist == nil {
		return err
	}
	
	return s.repo.Save(user)
}

func (s *UserService) DeleteUser(user *model.User) error {
	return s.repo.Delete(user)
}
