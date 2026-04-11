package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/models"
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*models.User, error)
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// Login validates credentials and returns a JWT token (simulated for now)
func (s *UserService) Login(email, password string) (string, error) {
	// Find user by email
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare the provided password with the hashed password in DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := "token"
	return token, nil
}
