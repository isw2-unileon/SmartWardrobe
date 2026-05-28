package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type ClothingItemRepository interface {
	GetAll() ([]models.ClothingItem, error)
	GetClothingItem(models.ClothingItem) ([]models.ClothingItem, error)
	AddClothingItem(models.ClothingItem) (*models.ClothingItem, error)
	UpdateClothingItem(int64, models.ClothingItem) (*models.ClothingItem, error)
	DeleteClothingItem(int64) error
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
			TypeId:   c.TypeId,
			ColorId:  c.ColorId,
			ImageUrl: c.ImageUrl,
			StyleId:  c.StyleId,
		})
	}

	return clothesDto, nil
}

func (s *ClothingItemService) GetClothingItem(clothingItem dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error) {
	model := models.ClothingItem{
		TypeId:   clothingItem.TypeId,
		ColorId:  clothingItem.ColorId,
		ImageUrl: clothingItem.ImageUrl,
		StyleId:  clothingItem.StyleId,
		UserId:   user.ID,
	}

	list, err := s.repo.GetClothingItem(model)
	if err != nil {
		return nil, err
	}

	// Convert the model to dto
	var listDto []dto.ClothingItemDto
	for _, c := range list {
		listDto = append(listDto, dto.ClothingItemDto{
			ID:       c.ID,
			TypeId:   c.TypeId,
			ColorId:  c.ColorId,
			ImageUrl: c.ImageUrl,
			StyleId:  c.StyleId,
		})
	}

	return listDto, nil
}

func (s *ClothingItemService) AddClothingItem(dto dto.ClothingItemDto, user dto.UserDto) (bool, error) {
	model := models.ClothingItem{
		TypeId:   dto.TypeId,
		ColorId:  dto.ColorId,
		ImageUrl: dto.ImageUrl,
		StyleId:  dto.StyleId,
		UserId:   user.ID,
	}

	save, err := s.repo.AddClothingItem(model)
	if err != nil {
		return false, err
	}

	if save.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (s *ClothingItemService) UpdateClothingItem(id int64, d dto.ClothingItemDto) (dto.ClothingItemDto, error) {
	model := models.ClothingItem{
		TypeId:   d.TypeId,
		ColorId:  d.ColorId,
		ImageUrl: d.ImageUrl,
		StyleId:  d.StyleId,
	}

	update, err := s.repo.UpdateClothingItem(id, model)

	//Convert the model to dto
	updateDto := dto.ClothingItemDto{
		ID:       update.ID,
		TypeId:   update.TypeId,
		ColorId:  update.ColorId,
		ImageUrl: update.ImageUrl,
		StyleId:  update.StyleId,
	}

	return updateDto, err
}

func (s *ClothingItemService) DeleteClothingItem(id int64) error {

	return s.repo.DeleteClothingItem(id)
}
