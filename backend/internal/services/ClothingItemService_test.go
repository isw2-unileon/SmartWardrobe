package services_test

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// The repository mock is created
type MockClothingItemRepository struct {
	mock.Mock
}

func (m *MockClothingItemRepository) GetClothingItemList(item models.ClothingItem) ([]models.ClothingItem, error) {
	args := m.Called(item)
	return args.Get(0).([]models.ClothingItem), args.Error(1)
}

func (m *MockClothingItemRepository) AddClothingItem(item models.ClothingItem) (*models.ClothingItem, error) {
	args := m.Called(item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ClothingItem), args.Error(1)
}

func (m *MockClothingItemRepository) UpdateClothingItem(id int64, item models.ClothingItem) (*models.ClothingItem, error) {
	args := m.Called(id, item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ClothingItem), args.Error(1)
}

func (m *MockClothingItemRepository) DeleteClothingItem(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestClothingItemService_GetAll_Success(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	fakeData := []models.ClothingItem{
		{
			ID:       1,
			ImageUrl: "url1",
			Type:     models.MasterType{ID: 1, Name: "Type1"},
			Color:    models.MasterColor{ID: 1, Name: "Color1"},
			Style:    models.MasterStyle{ID: 1, Name: "Style1"},
		},
	}

	mockRepo.On("GetClothingItemList", mock.Anything).Return(fakeData, nil)

	service := services.NewClothingItemService(mockRepo)

	result, err := service.GetAll(dto.UserDto{ID: "1"})

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "url1", result[0].ImageUrl)
	assert.Equal(t, int64(1), result[0].Type.ID)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_GetAll_Error(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	expectedError := errors.New("database error")
	var nilModels []models.ClothingItem

	mockRepo.On("GetClothingItemList", mock.Anything).Return(nilModels, expectedError)

	service := services.NewClothingItemService(mockRepo)

	result, err := service.GetAll(dto.UserDto{ID: "1"})

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_AddClothingItem_Success(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	// Datos de entrada
	inputDto := dto.ClothingItemDto{
		ImageUrl: "new_url",
		Type:     &dto.MasterTypeDto{ID: 1},
		Color:    &dto.MasterColorDto{ID: 2},
		Style:    &dto.MasterStyleDto{ID: 3},
	}
	inputUser := dto.UserDto{ID: "100"}

	savedItem := &models.ClothingItem{ID: 10}

	mockRepo.On("AddClothingItem", mock.Anything).Return(savedItem, nil)

	service := services.NewClothingItemService(mockRepo)

	success, err := service.AddClothingItem(inputDto, inputUser)

	assert.NoError(t, err)
	assert.True(t, success)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_AddClothingItem_Error(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	inputDto := dto.ClothingItemDto{
		Type:  &dto.MasterTypeDto{ID: 1},
		Color: &dto.MasterColorDto{ID: 2},
		Style: &dto.MasterStyleDto{ID: 3},
	}
	inputUser := dto.UserDto{ID: "100"}

	expectedError := errors.New("insert failed")

	mockRepo.On("AddClothingItem", mock.Anything).Return(nil, expectedError)

	service := services.NewClothingItemService(mockRepo)

	success, err := service.AddClothingItem(inputDto, inputUser)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.False(t, success)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_GetClothingItem_Success(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	inputDto := dto.ClothingItemDto{
		Type: &dto.MasterTypeDto{ID: 1},
	}
	inputUser := dto.UserDto{ID: "100"}

	fakeData := []models.ClothingItem{
		{
			ID:       5,
			ImageUrl: "url5",
			UserId:   "100",
			Type:     models.MasterType{ID: 1, Name: "Shirt"},
			Color:    models.MasterColor{ID: 2, Name: "Red"},
			Style:    models.MasterStyle{ID: 3, Name: "Casual"},
		},
	}

	mockRepo.On("GetClothingItemList", mock.Anything).Return(fakeData, nil)

	service := services.NewClothingItemService(mockRepo)
	result, err := service.GetClothingItem(inputDto, inputUser)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(5), result[0].ID)
	assert.Equal(t, "Shirt", result[0].Type.Name)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_GetClothingItem_Error(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	inputDto := dto.ClothingItemDto{}
	inputUser := dto.UserDto{ID: "100"}
	expectedError := errors.New("error fetching data")

	var nilModels []models.ClothingItem
	mockRepo.On("GetClothingItemList", mock.Anything).Return(nilModels, expectedError)

	service := services.NewClothingItemService(mockRepo)
	result, err := service.GetClothingItem(inputDto, inputUser)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_UpdateClothingItem_Success(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	id := int64(10)
	inputDto := dto.ClothingItemDto{
		ImageUrl: "updated_url",
		Type:     &dto.MasterTypeDto{ID: 2},
		Color:    &dto.MasterColorDto{ID: 3},
		Style:    &dto.MasterStyleDto{ID: 4},
	}

	updatedItem := &models.ClothingItem{
		ID:       10,
		ImageUrl: "updated_url",
		Type:     models.MasterType{ID: 2, Name: "Pants"},
		Color:    models.MasterColor{ID: 3, Name: "Blue"},
		Style:    models.MasterStyle{ID: 4, Name: "Formal"},
	}

	mockRepo.On("UpdateClothingItem", id, mock.Anything).Return(updatedItem, nil)

	service := services.NewClothingItemService(mockRepo)
	result, err := service.UpdateClothingItem(id, inputDto)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), result.ID)
	assert.Equal(t, "updated_url", result.ImageUrl)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_UpdateClothingItem_Error(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)

	id := int64(10)
	inputDto := dto.ClothingItemDto{
		ImageUrl: "updated_url",
		Type:     &dto.MasterTypeDto{ID: 2},
		Color:    &dto.MasterColorDto{ID: 3},
		Style:    &dto.MasterStyleDto{ID: 4},
	}
	expectedError := errors.New("update failed")

	mockRepo.On("UpdateClothingItem", id, mock.Anything).Return(nil, expectedError)

	service := services.NewClothingItemService(mockRepo)
	result, err := service.UpdateClothingItem(id, inputDto)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result.ID)

	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_DeleteClothingItem_Success(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)
	id := int64(99)

	mockRepo.On("DeleteClothingItem", id).Return(nil)

	service := services.NewClothingItemService(mockRepo)
	err := service.DeleteClothingItem(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestClothingItemService_DeleteClothingItem_Error(t *testing.T) {
	mockRepo := new(MockClothingItemRepository)
	id := int64(99)
	expectedError := errors.New("item not found")

	mockRepo.On("DeleteClothingItem", id).Return(expectedError)

	service := services.NewClothingItemService(mockRepo)
	err := service.DeleteClothingItem(id)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
