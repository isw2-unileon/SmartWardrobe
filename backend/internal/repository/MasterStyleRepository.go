package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// MasterStyleRepository handles database operations for the table master_styles
type MasterStyleRepository struct {
	db *gorm.DB
}

func NewMasterStyleRepository(db *gorm.DB) *MasterStyleRepository {
	return &MasterStyleRepository{db: db}
}

func (r *MasterStyleRepository) GetAll() ([]models.MasterStyle, error) {
	var styles []models.MasterStyle

	err := r.db.Find(&styles).Error

	return styles, err
}
