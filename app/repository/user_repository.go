package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transactgo/app/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Save(user *model.User) error
	Delete(user *model.User) error
}

type userRepository struct {
	users []model.User
}

func NewUserRepository() (UserRepository, error) {
	repo := &userRepository{}

	// Open the JSON file
	file, err := os.Open("data/users.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.users)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *userRepository) FindByUsername(username string) (*model.User,error) {
	for _, user := range r.users {
		if user.Username == username {
			return &user,nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) Save(user *model.User) error {
	// Update the user in the slice
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = *user
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(user *model.User) error {
	// Remove the user from the slice
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}
