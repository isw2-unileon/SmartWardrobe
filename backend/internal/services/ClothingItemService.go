package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type ClothingItemRepository interface {
	GetAll() ([]models.ClothingItem, error)
}

type ClothingItemService struct {
	repo ClothingItemRepository
}

func NewClothingItemService(repo ClothingItemRepository) *ClothingItemService {
	return &ClothingItemService{repo: repo}
}

// GetAll return all the clothes of the user
func (s *ClothingItemService) GetAll() ([]dto.ClothingItemDto, error) {
	clothes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert the model to dto
	var clothesDto []dto.ClothingItemDto
	for _, c := range clothes {
		clothesDto = append(clothesDto, dto.ClothingItemDto{
			ID:       c.ID,
			TypeId:   *c.TypeId,
			ColorId:  *c.ColorId,
			ImageUrl: c.ImageUrl,
			StyleId:  *c.StyleId,
		})
	}

	return clothesDto, nil

}
