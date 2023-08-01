package service

import (
	"errors"
	"transactgo/app/model"
	"transactgo/app/repository"

	"github.com/google/uuid"
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

func (s *UserService) UpdateUser(username string, userRequest *model.User) error {
	 user,err := s.repo.FindByUsername(username)
	if user == nil {
		return err
	}	

	user.Username = userRequest.Username
	user.Password = userRequest.Username

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(username string) error {
	user , err := s.repo.FindByUsername(username)
	if user == nil {
		return err
	}
	return s.repo.Delete(user)
}

func (s *UserService) AddUser (userRequest *model.User) error {
	user,err := s.repo.FindByUsername(userRequest.Username)
	if err != nil {
		// Return the error if there was an error when trying to find the user
		return err
	}
	if user != nil {
		// Return a custom error if a user with the same username already exists
		return errors.New("a user with this username already exists")
	}	
	userRequest.ID = uuid.New().String()
   return s.repo.Save(userRequest)
}