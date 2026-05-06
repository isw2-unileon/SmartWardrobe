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
type MockStyleService struct {
	mock.Mock
}

// Is assigned to behave himself
func (m *MockStyleService) GetAll() ([]dto.MasterStyleDto, error) {
	// m.Called() records that this function was called
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]dto.MasterStyleDto), args.Error(1)
}

func TestMasterStyleHandler_GetAll_Success(t *testing.T) {
	// The fake service is initialized
	mockService := new(MockStyleService)

	// Prepare the mock data to return
	fakeData := []dto.MasterStyleDto{
		{ID: 1, Name: "Casual"},
		{ID: 2, Name: "Formal"},
	}

	mockService.On("GetAll").Return(fakeData, nil)

	handler := handlers.NewMasterStyleHandler(mockService)

	// A fake HTTP environment is prepared with Gin
	router := gin.New()
	router.GET("/styles", handler.GetAll)

	// Create the false HTTP request and a recorder to catch the response
	req := httptest.NewRequest(http.MethodGet, "/styles", nil)
	recorder := httptest.NewRecorder()

	// The request is executed
	router.ServeHTTP(recorder, req)

	// The HTTP code is checked to be 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Is verified that the returned JSON is correct
	var responseBody []dto.MasterStyleDto
	err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.NoError(t, err)
	assert.Len(t, responseBody, 2)
	assert.Equal(t, "Casual", responseBody[0].Name)

	// Verified that the handler actually called the service
	mockService.AssertExpectations(t)
}

func TestMasterStyleHandler_GetAll_Error(t *testing.T) {
	// Prepare the mock error to return
	mockService := new(MockStyleService)
	mockService.On("GetAll").Return(nil, errors.New("error interno del servidor o de base de datos"))

	handler := handlers.NewMasterStyleHandler(mockService)

	router := gin.New()
	router.GET("/styles", handler.GetAll)

	// Create the false HTTP request and a recorder to catch the response
	req := httptest.NewRequest(http.MethodGet, "/styles", nil)
	recorder := httptest.NewRecorder()

	// The request is executed
	router.ServeHTTP(recorder, req)

	// The HTTP code is checked to be 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	// Is verified that the returned JSON is correct
	expectedJSON := `{"error":"Error when getting styles"}`
	assert.JSONEq(t, expectedJSON, recorder.Body.String())

	// Verified that the handler actually called the service
	mockService.AssertExpectations(t)
}
