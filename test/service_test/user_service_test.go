package service_test

import (
	"errors"
	"testing"
	"transactgo/app/model"
	"transactgo/app/service"
	"transactgo/test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserByUsername(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)

	result, err := mockService.GetUserByUsername("test")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Username, result.Username)

	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)
	mockRepo.On("Update", user).Return(nil)

	err := mockService.UpdateUser("test", user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)
	mockRepo.On("Delete", user).Return(nil)

	err := mockService.DeleteUser("test")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_AddUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(nil, errors.New("not found"))
	mockRepo.On("Save", user).Return(nil)

	err := mockService.AddUser(user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
func TestUserService_UpdateUser_Error(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(nil, errors.New("not found"))

	err := mockService.UpdateUser("test", user)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "not found")

	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_Error(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	mockRepo.On("FindByUsername", "test").Return(nil, errors.New("not found"))

	err := mockService.DeleteUser("test")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "not found")

	mockRepo.AssertExpectations(t)
}


func TestUserService_AddUser_Error(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)

	err := mockService.AddUser(user)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
