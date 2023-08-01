package service

import (
	"transactgo/app/model"
	"transactgo/app/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByUsername(username string) (*model.User,error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) UpdateUser(username string, user *model.User) error {
	 isExist,err := s.repo.FindByUsername(username)
	if isExist == nil {
		return err
	}	

	isExist.Username = user.Username
	isExist.Password = user.Password

	return s.repo.Save(isExist)
}

func (s *UserService) DeleteUser(username string) error {
	user , err := s.repo.FindByUsername(username)
	if user == nil {
		return err
	}
	return s.repo.Delete(user)
}
