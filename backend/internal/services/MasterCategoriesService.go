package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type MasterCategoriesRepository interface {
	GetByName(models.MasterCategory) (models.MasterCategory, error)
}

type MasterCategoriesService struct {
	repo MasterCategoriesRepository
}

func NewMasterCategoriesService(repo MasterCategoriesRepository) *MasterCategoriesService {
	return &MasterCategoriesService{repo: repo}
}

func (s *MasterCategoriesService) GetByName(name string) (*dto.MasterCategoryDto, error) {
	model := models.MasterCategory{
		Name: name,
	}
	category, err := s.repo.GetByName(model)
	if err != nil {
		return nil, err
	}

	// Convert the model to dto
	categoryDto := dto.MasterCategoryDto{
		ID:   category.ID,
		Name: category.Name,
	}

	return &categoryDto, nil
}
