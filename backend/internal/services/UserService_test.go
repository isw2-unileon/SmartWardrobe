package services_test

import (
	"errors"
	"group-15/backend/internal/models"
	"group-15/backend/internal/repository"
	"group-15/backend/internal/services"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	MockUser *models.User
	MockErr  error
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	return m.MockUser, m.MockErr
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := &MockUserRepository{
		MockUser: nil,
		MockErr:  errors.New("user not found in mock database"),
	}

	service := services.NewUserService(mockRepo)

	_, err := service.Login("dontExist", "123456")

	if err == nil {
		t.Errorf("We expected an error because the user doesn't exist, but it didn't give any")
	}
}

func TestLogin_Success(t *testing.T) {
	hashValido, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	hashedPassword := string(hashValido)

	mockRepo := &MockUserRepository{
		MockUser: &models.User{
			ID:       1,
			UserName: "test",
			Password: hashedPassword,
		},
		MockErr: nil,
	}

	service := services.NewUserService(mockRepo)

	token, err := service.Login("test", "123456")

	if err != nil {
		t.Errorf("We didn’t expect any mistakes, but it came out: %v", err)
	}
	if token == "" {
		t.Errorf("We were expecting a JWT token, but it returned an empty string")
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	repo := repository.NewUserRepository()
	service := services.NewUserService(repo)

	_, err := service.Login("test", "passWrong")

	if err == nil {
		t.Errorf("An error was expected due to incorrect password, but no error occurred")
	}
}
