package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type MasterStyleRepository interface {
	GetAll() ([]models.MasterStyle, error)
}

type MasterStyleService struct {
	repo MasterStyleRepository
}

func NewMasterStyleService(repo MasterStyleRepository) *MasterStyleService {
	return &MasterStyleService{repo: repo}
}

func (s *MasterStyleService) GetAll() ([]dto.MasterStyleDto, error) {
	styles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert the Model to DTO
	var styleDtos []dto.MasterStyleDto
	for _, c := range styles {
		styleDtos = append(styleDtos, dto.MasterStyleDto{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return styleDtos, nil
}
