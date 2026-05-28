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

// Add the clothing item
func (r *ClothingItemRepository) AddClothingItem(model models.ClothingItem) (*models.ClothingItem, error) {
	err := r.db.Create(&model).Error

	return &model, err
}

// Delete the clothing Item according to the id
func (r *ClothingItemRepository) DeleteClothingItem(id int64) error {
	err := r.db.Delete(&models.ClothingItem{}, id).Error
	return err
}
