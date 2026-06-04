package handlers_test

import (
	"backend/internal/dto"
	"backend/internal/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// The ClothingItemService mock is created
type MockClothingItemService struct {
	mock.Mock
}

// Implement the methods of the ClothingItemService interface
func (m *MockClothingItemService) GetByID(id int64) (dto.ClothingItemDto, error) {
	args := m.Called(id)

	var item dto.ClothingItemDto
	if args.Get(0) != nil {
		item = args.Get(0).(dto.ClothingItemDto)
	}

	return item, args.Error(1)
}

func (m *MockClothingItemService) GetAll(user dto.UserDto) ([]dto.ClothingItemDto, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.ClothingItemDto), args.Error(1)
}

func (m *MockClothingItemService) GetClothingItem(item dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error) {
	args := m.Called(item, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dto.ClothingItemDto), args.Error(1)
}

func (m *MockClothingItemService) AddClothingItem(item dto.ClothingItemDto, user dto.UserDto) (bool, error) {
	args := m.Called(item, user)
	return args.Bool(0), args.Error(1)
}

func (m *MockClothingItemService) UpdateClothingItem(id int64, item dto.ClothingItemDto) (dto.ClothingItemDto, error) {
	args := m.Called(id, item)
	return args.Get(0).(dto.ClothingItemDto), args.Error(1)
}

func (m *MockClothingItemService) DeleteClothingItem(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestClothingItemHandler_GetAll_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	fakeClothes := []dto.ClothingItemDto{
		{ID: 1, ImageUrl: "url1"},
	}
	expectedUser := dto.UserDto{ID: "user-123"}
	mockService.On("GetAll", expectedUser).Return(fakeClothes, nil)

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("userID", "user-123")
	})
	r.GET("/clothes", handler.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/clothes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []dto.ClothingItemDto
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, int64(1), response[0].ID)

	mockService.AssertExpectations(t)
}

func TestClothingItemHandler_GetAll_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	r := gin.New()

	r.GET("/clothes", handler.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/clothes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}

func TestClothingItemHandler_GetClothingItem_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	expectedUser := dto.UserDto{ID: "user-123"}
	expectedFilter := dto.ClothingItemDto{
		Type: &dto.MasterTypeDto{ID: 1},
	}

	mockService.On("GetClothingItem", expectedFilter, expectedUser).Return([]dto.ClothingItemDto{}, nil)

	r := gin.New()

	r.Use(func(c *gin.Context) {
		c.Set("userID", "user-123")
	})
	r.GET("/clothes/search", handler.GetClothingItem)

	req, _ := http.NewRequest(http.MethodGet, "/clothes/search?typeId=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestClothingItemHandler_GetClothingItem_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	r := gin.New()

	r.GET("/clothes/search", handler.GetClothingItem)

	req, _ := http.NewRequest(http.MethodGet, "/clothes/search", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}

func TestClothingItemHandler_AddClothingItem_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	inputDto := dto.ClothingItemDto{
		ImageUrl: "new-cloth.png",
		Type:     &dto.MasterTypeDto{ID: 1, Name: "T-shirt"},
		Color:    &dto.MasterColorDto{ID: 2, Name: "Blue"},
		Style:    &dto.MasterStyleDto{ID: 3, Name: "Casual"},
	}
	mockService.On("AddClothingItem", mock.Anything, dto.UserDto{ID: "user-123"}).Return(true, nil)

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("userID", "user-123")
	})
	r.POST("/clothes", handler.AddClothingItem)

	jsonBody, _ := json.Marshal(inputDto)
	req, _ := http.NewRequest(http.MethodPost, "/clothes", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "true", w.Body.String())
}

func TestClothingItemHandler_UpdateClothingItem_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	id := int64(45)

	bodyDto := dto.ClothingItemDto{
		ImageUrl: "updated.png",
		Type:     &dto.MasterTypeDto{ID: 1, Name: "T-shirt"},
		Color:    &dto.MasterColorDto{ID: 2, Name: "Blue"},
		Style:    &dto.MasterStyleDto{ID: 3, Name: "Casual"},
	}

	mockService.On("UpdateClothingItem", id, mock.Anything).Return(dto.ClothingItemDto{ID: id}, nil)

	r := gin.New()
	r.PUT("/clothes/:id", handler.UpdateClothingItem)

	jsonBody, _ := json.Marshal(bodyDto)
	req, _ := http.NewRequest(http.MethodPut, "/clothes/45", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestClothingItemHandler_UpdateClothingItem_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockClothingItemService)
	handler := handlers.NewClothingItemHandler(mockService)

	r := gin.New()
	r.PUT("/clothes/:id", handler.UpdateClothingItem)

	req, _ := http.NewRequest(http.MethodPut, "/clothes/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid ID")
}
