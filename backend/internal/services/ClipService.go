package services

import (
	"backend/internal/ai/clip"
	"backend/internal/dto"
	"io"
)

type ClipService struct {
	classifier *clip.CLIPClassifier
}

func NewClipService(classifier *clip.CLIPClassifier) *ClipService {
	return &ClipService{
		classifier: classifier,
	}
}

func (s *ClipService) Analyze(r io.Reader) (*dto.ClipPredictionResponse, error) {
	result, err := s.classifier.Classify(r)
	if err != nil {
		return nil, err
	}

	response := &dto.ClipPredictionResponse{
		Color: result.Color.Label,
		Style: result.Style.Label,
		Type:  result.Garment.Label,
	}

	return response, nil
}
