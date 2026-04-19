package services

import (
	"backend/internal/dto"
	"backend/internal/repository"
)

type MasterColorService struct {
	repo *repository.MasterColorRepository
}

func NewMasterColorService(repo *repository.MasterColorRepository) *MasterColorService {
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
