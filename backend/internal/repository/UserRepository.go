package repository

import (
	"backend/internal/models"
	"database/sql"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUserName searches for a user in the database
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := ""

	err := r.db.QueryRow(query, email).Scan(&user.UID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
