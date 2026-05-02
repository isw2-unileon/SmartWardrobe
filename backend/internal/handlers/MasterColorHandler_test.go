package handlers_test

import (
	"backend/internal/dto"
	"backend/internal/handlers"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// The service mock is created
type MockColorService struct {
	mock.Mock
}

// Is assigned to behave himself
func (m *MockColorService) GetAll() ([]dto.MasterColorDto, error) {
	// m.Called() records that this function was called
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]dto.MasterColorDto), args.Error(1)
}

func TestMasterColorHandler_GetAll_Success(t *testing.T) {
	// The fake service is initialized
	mockService := new(MockColorService)

	// Prepare the mock data to return
	fakeData := []dto.MasterColorDto{
		{ID: 1, Name: "Blue"},
		{ID: 2, Name: "White"},
	}

	mockService.On("GetAll").Return(fakeData, nil)

	handler := handlers.NewMasterColorHandler(mockService)

	// A fake HTTP environment is prepared with Gin
	router := gin.New()
	router.GET("/getAllColors", handler.GetAll)

	// Create the false HTTP request and a recorder to catch the response
	req := httptest.NewRequest(http.MethodGet, "/getAllColors", nil)
	recorder := httptest.NewRecorder()

	// The request is executed
	router.ServeHTTP(recorder, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Is verified that the returned JSON is correct
	var responseBody []dto.MasterColorDto
	err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.NoError(t, err)
	assert.Len(t, responseBody, 2)
	assert.Equal(t, "Blue", responseBody[0].Name)

	// Verified that the handler actually called the service
	mockService.AssertExpectations(t)
}

func TestMasterColorHandler_GetAll_Error(t *testing.T) {
	// Prepare the mock error to return
	mockService := new(MockColorService)
	mockService.On("GetAll").Return(nil, errors.New("error interno del servidor o de base de datos"))

	handler := handlers.NewMasterColorHandler(mockService)

	router := gin.New()
	router.GET("/getAllColors", handler.GetAll)

	// Create the false HTTP request and a recorder to catch the response
	req := httptest.NewRequest(http.MethodGet, "/getAllColors", nil)
	recorder := httptest.NewRecorder()

	// The request is executed
	router.ServeHTTP(recorder, req)

	// The HTTP code is checked to be 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	// Is verified that the returned JSON is correct
	expectedJSON := `{"error":"Error when getting colors"}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())

	// Verified that the handler actually called the service
	mockService.AssertExpectations(t)
}
