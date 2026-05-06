package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type MasterColorRepository interface {
	GetAll() ([]models.MasterColor, error)
}

type MasterColorService struct {
	repo MasterColorRepository
}

func NewMasterColorService(repo MasterColorRepository) *MasterColorService {
	return &MasterColorService{repo: repo}
}

func (s *MasterColorService) GetAll() ([]dto.MasterColorDto, error) {
	colors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert the Model to DTO
	var colorDtos []dto.MasterColorDto
	for _, c := range colors {
		colorDtos = append(colorDtos, dto.MasterColorDto{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return colorDtos, nil
}
