package services

import (
	"backend/internal/dto"
	"backend/internal/models"
)

type MasterTypeRepository interface {
	GetAll() ([]models.MasterType, error)
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
