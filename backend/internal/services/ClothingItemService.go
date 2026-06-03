package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type ClothingItemRepository interface {
	GetAll() ([]models.ClothingItem, error)
	GetClothingItem(models.ClothingItem) ([]models.ClothingItem, error)
	GetByID(int64) (*models.ClothingItem, error)
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
		clothesDto = append(clothesDto, mapModelToDto(c))
	}

	return clothesDto, nil
}

func (s *ClothingItemService) GetClothingItem(clothingItem dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error) {
	model := models.ClothingItem{
		UserId: user.ID,
	}

	// Only assigns the filter if the field is not nil
	if clothingItem.Type != nil {
		model.TypeId = &clothingItem.Type.ID
	}
	if clothingItem.Color != nil {
		model.ColorId = &clothingItem.Color.ID
	}
	if clothingItem.Style != nil {
		model.StyleId = &clothingItem.Style.ID
	}

	list, err := s.repo.GetClothingItem(model)
	if err != nil {
		return nil, err
	}

	// Convert the model to dto
	var listDto []dto.ClothingItemDto
	for _, c := range list {
		listDto = append(listDto, mapModelToDto(c))
	}

	return listDto, nil
}

// Get the clothing item in function of filters
func (s *ClothingItemService) GetByID(id int64) (dto.ClothingItemDto, error) {
	model, err := s.repo.GetByID(id)
	if err != nil {
		return dto.ClothingItemDto{}, err
	}

	return mapModelToDto(*model), nil
}

func (s *ClothingItemService) AddClothingItem(dto dto.ClothingItemDto, user dto.UserDto) (bool, error) {
	model := models.ClothingItem{
		TypeId:   &dto.Type.ID,
		ColorId:  &dto.Color.ID,
		ImageUrl: dto.ImageUrl,
		StyleId:  &dto.Style.ID,
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
		TypeId:   &d.Type.ID,
		ColorId:  &d.Color.ID,
		ImageUrl: d.ImageUrl,
		StyleId:  &d.Style.ID,
	}

	update, err := s.repo.UpdateClothingItem(id, model)
	if err != nil || update == nil {
		return dto.ClothingItemDto{}, err
	}

	//Convert the model to dto
	updateDto := mapModelToDto(*update)

	return updateDto, err
}

func (s *ClothingItemService) DeleteClothingItem(id int64) error {

	return s.repo.DeleteClothingItem(id)
}

// Convert the model to dto
func mapModelToDto(c models.ClothingItem) dto.ClothingItemDto {
	return dto.ClothingItemDto{
		ID:       c.ID,
		ImageUrl: c.ImageUrl,
		Type: &dto.MasterTypeDto{
			ID:   c.Type.ID,
			Name: c.Type.Name,
		},
		Color: &dto.MasterColorDto{
			ID:   c.Color.ID,
			Name: c.Color.Name,
		},
		Style: &dto.MasterStyleDto{
			ID:   c.Style.ID,
			Name: c.Style.Name,
		},
	}
}
