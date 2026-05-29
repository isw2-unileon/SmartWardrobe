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

	err := r.db.
		Preload("Type").
		Preload("Color").
		Preload("Style").
		Find(&clothes).Error

	return clothes, err
}

func (r *ClothingItemRepository) GetClothingItem(filters models.ClothingItem) ([]models.ClothingItem, error) {
	var list []models.ClothingItem

	err := r.db.
		Preload("Type").
		Preload("Color").
		Preload("Style").
		Where(filters).
		Find(&list).Error

	return list, err
}

// Add the clothing item
func (r *ClothingItemRepository) AddClothingItem(model models.ClothingItem) (*models.ClothingItem, error) {
	err := r.db.Create(&model).Error

	return &model, err
}

// Update the clothing item according to the id with the params
func (r *ClothingItemRepository) UpdateClothingItem(id int64, model models.ClothingItem) (*models.ClothingItem, error) {
	err := r.db.
		Model(&models.ClothingItem{}).
		Where("id = ?", id).
		Updates(model).Error

	if err != nil {
		return nil, err
	}

	var updated models.ClothingItem
	err = r.db.
		Preload("Type").
		Preload("Color").
		Preload("Style").
		First(&updated, id).Error
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete the clothing Item according to the id
func (r *ClothingItemRepository) DeleteClothingItem(id int64) error {
	err := r.db.Delete(&models.ClothingItem{}, id).Error
	return err
}
