package services

import (
	"backend/internal/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type LocationService struct {
	baseURL string
}

func NewLocationService() *LocationService {
	return &LocationService{baseURL: "https://geocoding-api.open-meteo.com/v1/search"}
}

// Used only for testing
func NewLocationServiceWithURL(baseURL string) *LocationService {
	return &LocationService{baseURL: baseURL}
}

func (s *LocationService) GetLocation(city string, country string) (*dto.LocationDto, error) {
	// Obtain the coordinates of the city
	geoURL := fmt.Sprintf("%s?name=%s&count=10&format=json", s.baseURL, url.QueryEscape(city))

	resp, err := http.Get(geoURL)
	if err != nil {
		fmt.Printf("error search the city: %v\n", err)
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)

	var location dto.LocationDto
	if err := json.Unmarshal(body, &location); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	if len(location.Results) == 0 {
		fmt.Printf("not found the city: %s\n", city)
		return nil, err
	}

	// Filter the list of results to find the correct country
	for _, result := range location.Results {
		if strings.EqualFold(result.Country, country) {
			// dto with the correct location
			resultadoFiltrado := dto.LocationDto{
				Results: []struct {
					Name      string  `json:"name" binding:"required"`
					Country   string  `json:"country" binding:"required"`
					Latitude  float64 `json:"latitude" binding:"required"`
					Longitude float64 `json:"longitude" binding:"required"`
				}{result},
			}
			return &resultadoFiltrado, nil
		}
	}

	return nil, fmt.Errorf("the city %s was found, but not in the country %s", city, country)
}
