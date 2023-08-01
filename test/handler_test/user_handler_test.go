package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactgo/app/handler"
	"transactgo/app/model"
	"transactgo/app/service"
	"transactgo/test/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_GetUserByUsername(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler.NewUserHandler(mockService, router)

	req, _ := http.NewRequest(http.MethodGet, "/users/test", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	mockRepo.AssertExpectations(t)
}

func TestUserHandler_UpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
		Password: "password",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)
	mockRepo.On("Update", user).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler.NewUserHandler(mockService, router)

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/users/test", bytes.NewBuffer(jsonUser))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	mockRepo.AssertExpectations(t)
}

func TestUserHandler_DeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockService := service.NewUserService(mockRepo)

	user := &model.User{
		Username: "test",
	}

	mockRepo.On("FindByUsername", "test").Return(user, nil)
	mockRepo.On("Delete", user).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler.NewUserHandler(mockService, router)

	req, _ := http.NewRequest(http.MethodDelete, "/users/test", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	mockRepo.AssertExpectations(t)
}

func TestUserHandler_AddUser(t *testing.T) {
    mockRepo := new(mocks.UserRepository)
    mockService := service.NewUserService(mockRepo)

    user := &model.User{
        Username: "test",
        Password: "password",
    }

    mockRepo.On("FindByUsername", "test").Return(nil, nil)
    mockRepo.On("Save", mock.Anything).Return(nil) // Gunakan mock.Anything di sini

    gin.SetMode(gin.TestMode)
    router := gin.Default()
    handler.NewUserHandler(mockService, router)

    jsonUser, _ := json.Marshal(user)
    req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonUser))
    resp := httptest.NewRecorder()

    router.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)

    mockRepo.AssertExpectations(t)
}


