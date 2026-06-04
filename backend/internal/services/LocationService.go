package services

import (
	"backend/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type LocationService struct{}

func NewLocationService() *LocationService {
	return &LocationService{}
}

func (s *LocationService) GetLocation(city string) (*dto.LocationDto, error) {
	// Obtain the coordinates of the city
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&format=json", url.QueryEscape(city))

	respGeo, err := http.Get(geoURL)
	if err != nil {
		fmt.Printf("error search the city: %v\n", err)
		return nil, err
	}
	defer respGeo.Body.Close()

	var location dto.LocationDto
	json.NewDecoder(respGeo.Body).Decode(&location)

	if len(location.Results) == 0 {
		fmt.Printf("not found the city: %s\n", city)
		return nil, err
	}

	return &location, nil
}
