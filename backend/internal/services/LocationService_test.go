package services_test

import (
	"backend/internal/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock
func newMockLocationServer(statusCode int, body interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
}

func fakeLocationResponse(results []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"results": results,
	}
}

func TestLocationService_GetLocation_Success(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := fakeLocationResponse([]map[string]interface{}{
		{"name": "Madrid", "country": "Spain", "latitude": 40.4168, "longitude": -3.7038},
	})

	// The fake repository is initialized
	server := newMockLocationServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewLocationServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetLocation("Madrid", "Spain")

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "Madrid", result.Results[0].Name)
	assert.Equal(t, "Spain", result.Results[0].Country)
	assert.Equal(t, 40.4168, result.Results[0].Latitude)
	assert.Equal(t, -3.7038, result.Results[0].Longitude)
}

func TestLocationService_GetLocation_FiltersByCountry(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := fakeLocationResponse([]map[string]interface{}{
		{"name": "Leon", "country": "Mexico", "latitude": 21.1167, "longitude": -101.6833},
		{"name": "Leon", "country": "Spain", "latitude": 42.5987, "longitude": -5.5671},
		{"name": "Leon", "country": "Nicaragua", "latitude": 12.4333, "longitude": -86.8833},
	})

	// The fake repository is initialized
	server := newMockLocationServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewLocationServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetLocation("Leon", "Spain")

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "Spain", result.Results[0].Country)
	assert.Equal(t, 42.5987, result.Results[0].Latitude)
}

func TestLocationService_GetLocation_CountryNotFound(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := fakeLocationResponse([]map[string]interface{}{
		{"name": "Madrid", "country": "Mexico", "latitude": 22.0, "longitude": -100.0},
	})

	// The fake repository is initialized
	server := newMockLocationServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewLocationServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetLocation("Madrid", "Spain")

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Madrid")
	assert.Contains(t, err.Error(), "Spain")
}

func TestLocationService_GetLocation_CityNotFound(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := map[string]interface{}{
		"results": []map[string]interface{}{},
	}

	// The fake repository is initialized
	server := newMockLocationServer(http.StatusOK, fakeResponse)
	defer server.Close()

	svc := services.NewLocationServiceWithURL(server.URL)
	// The function is executed
	result, _ := svc.GetLocation("CiudadInexistente", "Spain")

	// The results are checked
	assert.Nil(t, result)
}

func TestLocationService_GetLocation_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("not valid json {{{"))
	}))
	defer server.Close()

	svc := services.NewLocationServiceWithURL(server.URL)
	// The function is executed
	result, err := svc.GetLocation("Madrid", "Spain")

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLocationService_GetLocation_CaseInsensitiveCountry(t *testing.T) {
	// Prepare the mock data to return
	fakeResponse := fakeLocationResponse([]map[string]interface{}{
		{"name": "Madrid", "country": "Spain", "latitude": 40.4168, "longitude": -3.7038},
	})

	// The fake repository is initialized
	server := newMockLocationServer(http.StatusOK, fakeResponse)
	defer server.Close()

	// The function is executed
	svc := services.NewLocationServiceWithURL(server.URL)

	// The results are checked
	result, err := svc.GetLocation("Madrid", "spain")
	assert.NoError(t, err)
	assert.NotNil(t, result)

	result, err = svc.GetLocation("Madrid", "SPAIN")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
