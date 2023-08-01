package service

import (
	"errors"
	"transactgo/app/model"
	"transactgo/app/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	user,_ := s.repo.FindByUsername(userRequest.Username)
	if user != nil {
		// Return a custom error if a user with the same username already exists
		return errors.New("a user with this username already exists")
	}	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
    }
	userRequest.ID = uuid.New().String()
	userRequest.Password = string(hashedPassword)
   return s.repo.Save(userRequest)
}