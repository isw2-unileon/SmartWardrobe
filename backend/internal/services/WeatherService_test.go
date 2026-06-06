package services_test

import (
	"backend/internal/dto"
	"backend/internal/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock
func newMockWeatherServer(statusCode int, body interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
}

func defaultLocation() *dto.LocationDto {
	return &dto.LocationDto{
		Results: []struct {
			Name      string  `json:"name" binding:"required"`
			Country   string  `json:"country" binding:"required"`
			Latitude  float64 `json:"latitude" binding:"required"`
			Longitude float64 `json:"longitude" binding:"required"`
		}{
			{Latitude: 40.4168, Longitude: -3.7038},
		},
	}
}

func TestWeatherService_GetWeather_Success(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := map[string]interface{}{
		"daily": map[string]interface{}{
			"time":               []string{"2026-06-04", "2026-06-05"},
			"temperature_2m_max": []float64{25.0, 22.0},
			"temperature_2m_min": []float64{15.0, 13.0},
		},
	}

	// The fake repository is initialized
	server := newMockWeatherServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewWeatherServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetWeather(defaultLocation(), "2026-06-04", "2026-06-05")

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "2026-06-04", result[0].Date)
	assert.Equal(t, 25.0, *result[0].MaxTemp)
	assert.Equal(t, 15.0, *result[0].MinTemp)
	assert.Equal(t, "2026-06-05", result[1].Date)
	assert.Equal(t, 22.0, *result[1].MaxTemp)
	assert.Equal(t, 13.0, *result[1].MinTemp)
}

func TestWeatherService_GetWeather_SingleDay(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := map[string]interface{}{
		"daily": map[string]interface{}{
			"time":               []string{"2026-06-04"},
			"temperature_2m_max": []float64{30.0},
			"temperature_2m_min": []float64{20.0},
		},
	}

	// The fake repository is initialized
	server := newMockWeatherServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewWeatherServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetWeather(defaultLocation(), "2026-06-04", "2026-06-04")

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "2026-06-04", result[0].Date)
}

func TestWeatherService_GetWeather_EmptyData(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := map[string]interface{}{
		"daily": map[string]interface{}{
			"time":               []string{},
			"temperature_2m_max": []float64{},
			"temperature_2m_min": []float64{},
		},
	}

	// The fake repository is initialized
	server := newMockWeatherServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewWeatherServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetWeather(defaultLocation(), "2026-06-04", "2026-06-04")

	// The results are checked
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestWeatherService_GetWeather_APIError(t *testing.T) {
	// The fake repository is initialized
	server := newMockWeatherServer(http.StatusInternalServerError, nil)
	defer server.Close()

	svc := services.NewWeatherServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetWeather(defaultLocation(), "2026-06-04", "2026-06-04")

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "500")
}

func TestWeatherService_GetWeather_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("not valid json {{{"))
	}))
	defer server.Close()

	svc := services.NewWeatherServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetWeather(defaultLocation(), "2026-06-04", "2026-06-04")

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
}
