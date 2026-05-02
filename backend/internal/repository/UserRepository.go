package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUserName searches for a user in the database
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := r.db.Find(&user).Error

	return user, err
}
