package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"group-15/backend/internal/models"
)

type UserRepositoryInterface interface {
	FindByUsername(userName string) (*models.User, error)
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// Login validates credentials and returns a JWT token (simulated for now)
func (s *UserService) Login(userName, password string) (string, error) {
	// Find user by username
	user, err := s.repo.FindByUsername(userName)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare the provided password with the hashed password in DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Get the remote server key
	secretString := os.Getenv("JWT_SECRET")

	if secretString == "" {
		return "", errors.New("Internal Error Server")
	}

	jwtSecretKey := []byte(secretString)

	// Generate JWT Token
	claims := jwt.MapClaims{
		"sub":      user.ID,
		"userName": user.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", errors.New("error al generar el token")
	}
	return tokenString, nil
}
