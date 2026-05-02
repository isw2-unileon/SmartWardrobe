package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// MasterTypeRepository handles database operations for the table master_types
type MasterTypeRepository struct {
	db *gorm.DB
}

func NewMasterTypeRepository(db *gorm.DB) *MasterTypeRepository {
	return &MasterTypeRepository{db: db}
}

func (r *MasterTypeRepository) GetAll() ([]models.MasterType, error) {
	var types []models.MasterType

	err := r.db.Find(&types).Error

	return types, err
}
