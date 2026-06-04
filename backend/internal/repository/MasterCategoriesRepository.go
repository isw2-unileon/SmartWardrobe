package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// MasterCategoriesRepository handles database operations for the table master_types
type MasterCategoriesRepository struct {
	db *gorm.DB
}

func NewMasterCategoriesRepository(db *gorm.DB) *MasterCategoriesRepository {
	return &MasterCategoriesRepository{db: db}
}

func (r *MasterCategoriesRepository) GetByName(name models.MasterCategory) (models.MasterCategory, error) {
	var category models.MasterCategory

	err := r.db.Where(name).Find(&category).Error

	return category, err
}
