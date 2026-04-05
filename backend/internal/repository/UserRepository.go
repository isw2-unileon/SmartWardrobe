package repository

import (
	"errors"
	"group-15/backend/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	// Here you would inject your database connection (e.g., db *sql.DB)
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// FindByUserName searches for a user in the database
func (r *UserRepository) FindByUsername(userName string) (*models.User, error) {
	// MOCK DATA: Simulating a database search for testing
	if userName == "test" {
		return &models.User{
			ID:       1,
			UserName: "test",
			// This is a hashed version of "123456" using bcrypt
			Password: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
		}, nil
	}

	return nil, errors.New("user not found")
}
