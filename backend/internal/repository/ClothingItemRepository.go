package repository

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

// ClothingItemRepository handles database operations for the table clothing_items
type ClothingItemRepository struct {
	db *gorm.DB
}

func NewClothingItemRepository(db *gorm.DB) *ClothingItemRepository {
	return &ClothingItemRepository{db: db}
}

func (r *ClothingItemRepository) GetAll() ([]models.ClothingItem, error) {
	var clothes []models.ClothingItem

	err := r.db.Find(&clothes).Error

	return clothes, err
}
