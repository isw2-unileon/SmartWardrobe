package services

import (
	"backend/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeather(city *dto.LocationDto, startDate string, endDate string) (*dto.WeatherDto, error) {
	baseURL := "https://api.open-meteo.com/v1/forecast"

	// The url is build param for param
	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%.4f", city.Results[0].Latitude))
	params.Add("longitude", fmt.Sprintf("%.4f", city.Results[0].Longitude))
	params.Add("daily", "temperature_2m_max,temperature_2m_min")
	params.Add("timezone", "auto")
	params.Add("start_date", startDate)
	params.Add("end_date", endDate)

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var weather dto.WeatherDto
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	if len(weather.Daily.Time) > 0 {
		date := weather.Daily.Time[0]
		max := weather.Daily.MaxTemp[0]
		min := weather.Daily.MinTemp[0]

		fmt.Printf("The weather for the day %s is: Max %.1f°C / Min %.1f°C\n", date, max, min)
	} else {
		fmt.Println("No data were found for that date.")
	}

	return &weather, nil

}
