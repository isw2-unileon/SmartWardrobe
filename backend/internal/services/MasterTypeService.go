package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type MasterTypeRepository interface {
	GetAll() ([]models.MasterType, error)
	GetTypesByCategory(models.MasterType) ([]models.MasterType, error)
}

type MasterTypeService struct {
	repo MasterTypeRepository
}

func NewMasterTypeService(repo MasterTypeRepository) *MasterTypeService {
	return &MasterTypeService{repo: repo}
}

func (s *MasterTypeService) GetAll() ([]dto.MasterTypeDto, error) {
	types, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert the Model to DTO
	var typeDtos []dto.MasterTypeDto
	for _, c := range types {
		typeDtos = append(typeDtos, dto.MasterTypeDto{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return typeDtos, nil
}

func (s *MasterTypeService) GetTypesWithTempRangeAndCategory(weather dto.WeatherDayDto, category dto.MasterCategoryDto) ([]dto.MasterTypeDto, error) {
	search := models.MasterType{
		CategoryId: &category.ID,
	}
	types, err := s.repo.GetTypesByCategory(search)
	if err != nil {
		return nil, err
	}

	// Convert the Model to DTO if apply the conditions
	var typeDtos []dto.MasterTypeDto
	for _, c := range types {
		if (c.MinTemp == nil && c.MaxTemp == nil) ||
			(c.MaxTemp != nil && (*c.MaxTemp) >= *weather.MaxTemp) &&
				(c.MinTemp != nil && (*c.MinTemp) >= *weather.MinTemp) ||
			(c.MinTemp == nil && c.MaxTemp != nil && (*c.MaxTemp) >= *weather.MaxTemp) ||
			(c.MaxTemp == nil && c.MinTemp != nil && (*c.MinTemp) >= *weather.MinTemp) {
			typeDtos = append(typeDtos, dto.MasterTypeDto{
				ID:   c.ID,
				Name: c.Name,
				Category: &dto.MasterCategoryDto{
					ID: *c.CategoryId,
				},
			})

		}
	}

	return typeDtos, nil
}
