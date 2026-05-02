package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// MasterColorRepository handles database operations for the table master_colors
type MasterColorRepository struct {
	db *gorm.DB
}

func NewMasterColorRepository(db *gorm.DB) *MasterColorRepository {
	return &MasterColorRepository{db: db}
}

func (r *MasterColorRepository) GetAll() ([]models.MasterColor, error) {
	var colors []models.MasterColor

	err := r.db.Find(&colors).Error

	return colors, err
}
