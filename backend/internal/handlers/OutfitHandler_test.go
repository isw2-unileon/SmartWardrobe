package handlers_test

import (
	"backend/internal/dto"
	"backend/internal/handlers"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock
type MockOutfitService struct {
	mock.Mock
}

func (m *MockOutfitService) GenerateOutfit(req dto.OutfitRequestDto, user dto.UserDto) ([]dto.OutfitResponseDto, error) {
	args := m.Called(req, user)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]dto.OutfitResponseDto), args.Error(1)
}

// Helpers
func setupOutfitRouter(svc *MockOutfitService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handlers.NewOutfitHandler(svc)

	r.Use(func(c *gin.Context) {
		c.Set("userID", "user-uuid-123")
		c.Next()
	})

	r.POST("/api/generateOutfit", h.GenerateOutfit)
	return r
}

func setupOutfitRouterNoAuth(svc *MockOutfitService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handlers.NewOutfitHandler(svc)

	r.POST("/api/generateOutfit", h.GenerateOutfit)
	return r
}

func outfitRequestBody(req dto.OutfitRequestDto) *bytes.Buffer {
	body, _ := json.Marshal(req)
	return bytes.NewBuffer(body)
}

func defaultOutfitRequest() dto.OutfitRequestDto {
	return dto.OutfitRequestDto{
		City:      "Madrid",
		Country:   "Spain",
		StartDate: "2026-06-04",
		EndDate:   "2026-06-04",
	}
}

func defaultOutfitResponse() []dto.OutfitResponseDto {
	maxTemp := 25.0
	minTemp := 15.0
	return []dto.OutfitResponseDto{
		{
			Weather: dto.WeatherDayDto{Date: "2026-06-04", MaxTemp: &maxTemp, MinTemp: &minTemp},
			Outfit: dto.OutfitDto{
				Upperwear:  &dto.ClothingItemDto{ImageUrl: "https://img.com/upper.jpg"},
				Bottomwear: &dto.ClothingItemDto{ImageUrl: "https://img.com/bottom.jpg"},
				Footwear:   &dto.ClothingItemDto{ImageUrl: "https://img.com/foot.jpg"},
				Outerwear:  nil,
			},
		},
	}
}

func TestOutfitHandler_GenerateOutfit_Success(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)
	mockSvc.On("GenerateOutfit", mock.Anything, mock.Anything).
		Return(defaultOutfitResponse(), nil)

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(defaultOutfitRequest()))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	var result []dto.OutfitResponseDto
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "2026-06-04", result[0].Weather.Date)
	// Verified that the handler actually called the service
	mockSvc.AssertExpectations(t)
}

func TestOutfitHandler_GenerateOutfit_Unauthorized(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)

	r := setupOutfitRouterNoAuth(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(defaultOutfitRequest()))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 401 UNAUTHORIZED
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "Unauthorized", body["error"])
	// Verified that the handler actually not called the service
	mockSvc.AssertNotCalled(t, "GenerateOutfit")
}

func TestOutfitHandler_GenerateOutfit_InvalidBody(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		bytes.NewBufferString("not valid json {{{"))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 400 BAD REQUEST
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Verified that the handler actually not called the service
	mockSvc.AssertNotCalled(t, "GenerateOutfit")
}

func TestOutfitHandler_GenerateOutfit_ServiceError(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)
	mockSvc.On("GenerateOutfit", mock.Anything, mock.Anything).
		Return(nil, errors.New("weather API error"))

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(defaultOutfitRequest()))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 500 INTERNAL SERVER ERROR
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var body map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "Error when filter clothes.", body["error"])
	// Verified that the handler actually called the service
	mockSvc.AssertExpectations(t)
}

func TestOutfitHandler_GenerateOutfit_PassesUserToService(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)

	// Verifica que el userID del contexto llega correctamente al service
	expectedUser := dto.UserDto{ID: "user-uuid-123"}
	mockSvc.On("GenerateOutfit", mock.Anything, expectedUser).
		Return(defaultOutfitResponse(), nil)

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(defaultOutfitRequest()))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
	// Verified that the handler actually called the service
	mockSvc.AssertExpectations(t)
}

func TestOutfitHandler_GenerateOutfit_PassesRequestToService(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)

	expectedReq := defaultOutfitRequest()
	mockSvc.On("GenerateOutfit", expectedReq, mock.Anything).
		Return(defaultOutfitResponse(), nil)

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(expectedReq))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
	// Verified that the handler actually called the service
	mockSvc.AssertExpectations(t)
}

func TestOutfitHandler_GenerateOutfit_MultipleDays(t *testing.T) {
	// The fake service is initialized
	mockSvc := new(MockOutfitService)

	// Prepare the mock data to return
	maxTemp1, minTemp1 := 25.0, 15.0
	maxTemp2, minTemp2 := 20.0, 12.0
	multiDayResponse := []dto.OutfitResponseDto{
		{Weather: dto.WeatherDayDto{Date: "2026-06-04", MaxTemp: &maxTemp1, MinTemp: &minTemp1}},
		{Weather: dto.WeatherDayDto{Date: "2026-06-05", MaxTemp: &maxTemp2, MinTemp: &minTemp2}},
	}
	mockSvc.On("GenerateOutfit", mock.Anything, mock.Anything).
		Return(multiDayResponse, nil)

	r := setupOutfitRouter(mockSvc)
	w := httptest.NewRecorder()

	// Create the false HTTP request and a recorder to catch the response
	req, _ := http.NewRequest(http.MethodPost, "/api/generateOutfit",
		outfitRequestBody(dto.OutfitRequestDto{
			City:      "Madrid",
			Country:   "Spain",
			StartDate: "2026-06-04",
			EndDate:   "2026-06-05",
		}))
	req.Header.Set("Content-Type", "application/json")

	// The request is executed
	r.ServeHTTP(w, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	var result []dto.OutfitResponseDto
	_ = json.Unmarshal(w.Body.Bytes(), &result)
	assert.Len(t, result, 2)

	// Verified that the handler actually called the service
	mockSvc.AssertExpectations(t)
}
