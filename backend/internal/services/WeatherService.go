package services

import (
	"backend/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeather(city string) (*dto.WeatherDto, error) {
	url := fmt.Sprintf("https://wttr.in/%s?format=j1&lang=es", city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error calling wttr.in: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wttr.in returned status: %d", resp.StatusCode)
	}

	var raw struct {
		Weather []struct {
			MinTempC string `json:"mintempC"`
			MaxTempC string `json:"maxtempC"`
		} `json:"weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	if len(raw.Weather) == 0 {
		return nil, fmt.Errorf("no weather data for city: %s", city)
	}

	// Pass the string to float64
	var minTemp, maxTemp float64
	fmt.Sscanf(raw.Weather[0].MinTempC, "%f", &minTemp)
	fmt.Sscanf(raw.Weather[0].MaxTempC, "%f", &maxTemp)

	return &dto.WeatherDto{
		MinTemp: minTemp,
		MaxTemp: maxTemp,
		City:    city,
	}, nil
}
