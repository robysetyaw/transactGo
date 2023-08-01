package repository_test

import (
	"errors"
	"testing"
	"transactgo/app/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	user, _ := args.Get(0).(*model.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Save(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestFindByUsername(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)

	result, err := mockRepo.FindByUsername("test")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Username, result.Username)

	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Update", user).Return(nil)

	err := mockRepo.Update(user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Delete", user).Return(nil)

	err := mockRepo.Delete(user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestSave(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Save", user).Return(nil)

	err := mockRepo.Save(user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFindByUsername_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("FindByUsername", "test").Return(nil, errors.New("Error occured"))

	_, err := mockRepo.FindByUsername("test")

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdate_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Update", user).Return(errors.New("Error occured"))

	err := mockRepo.Update(user)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDelete_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Delete", user).Return(errors.New("Error occured"))

	err := mockRepo.Delete(user)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestSave_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("Save", user).Return(errors.New("Error occured"))

	err := mockRepo.Save(user)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
