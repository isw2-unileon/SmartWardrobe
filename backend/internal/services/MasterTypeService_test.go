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
type MockTypeRepository struct {
	mock.Mock
}

// Is assigned to behave himself
func (m *MockTypeRepository) GetAll() ([]models.MasterType, error) {
	// m.Called() records that this function was called
	args := m.Called()

	return args.Get(0).([]models.MasterType), args.Error(1)
}

func (m *MockTypeRepository) GetTypesByCategory(models.MasterType) ([]models.MasterType, error) {
	// m.Called() records that this function was called
	args := m.Called()

	return args.Get(0).([]models.MasterType), args.Error(1)
}

// Helpers

func floatPtr(v float64) *float64 { return &v }
func int64Ptr(v int64) *int64     { return &v }

func TestMasterTypeService_GetAll_Success(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "T-shirt"},
		{ID: 2, Name: "Hoodie"},
	}

	mockRepo.On("GetAll").Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)

	// The function is executed
	result, err := service.GetAll()

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(1), result[0].ID)
	assert.Equal(t, "T-shirt", result[0].Name)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetAll_Error(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock error to return
	expectedError := errors.New("error fatal: database disconected")

	var nilModels []models.MasterType
	mockRepo.On("GetAll").Return(nilModels, expectedError)

	service := services.NewMasterTypeService(mockRepo)

	// The function is executed
	result, err := service.GetAll()

	// The results are checked
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetAll_Empty(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	mockRepo.On("GetAll").Return([]models.MasterType{}, nil)

	service := services.NewMasterTypeService(mockRepo)

	// The function is executed
	result, err := service.GetAll()

	// The results are checked
	assert.NoError(t, err)
	assert.Empty(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_RepoError(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock error to return
	expectedError := errors.New("db error")
	var nilModels []models.MasterType
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(nilModels, expectedError)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(25.0), MinTemp: floatPtr(15.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_NoTempRestriction(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Camiseta", CategoryId: int64Ptr(1), MinTemp: nil, MaxTemp: nil},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(30.0), MinTemp: floatPtr(20.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Camiseta", result[0].Name)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_BothTempMatch(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Abrigo", CategoryId: int64Ptr(1), MinTemp: floatPtr(5.0), MaxTemp: floatPtr(15.0)},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(25.0), MinTemp: floatPtr(18.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Abrigo", result[0].Name)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_BothTempNoMatch(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Abrigo", CategoryId: int64Ptr(1), MinTemp: floatPtr(20.0), MaxTemp: floatPtr(30.0)},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(10.0), MinTemp: floatPtr(5.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Empty(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_OnlyMaxTemp(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Camiseta", CategoryId: int64Ptr(1), MinTemp: nil, MaxTemp: floatPtr(15.0)},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(25.0), MinTemp: floatPtr(18.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_OnlyMinTemp(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Camiseta", CategoryId: int64Ptr(1), MinTemp: floatPtr(10.0), MaxTemp: nil},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(25.0), MinTemp: floatPtr(18.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_MultipleTypes(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Camiseta", CategoryId: int64Ptr(1), MinTemp: nil, MaxTemp: nil},
		{ID: 2, Name: "Abrigo", CategoryId: int64Ptr(1), MinTemp: floatPtr(5.0), MaxTemp: floatPtr(10.0)},
		{ID: 3, Name: "Plumífero", CategoryId: int64Ptr(1), MinTemp: floatPtr(25.0), MaxTemp: floatPtr(30.0)},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(20.0), MinTemp: floatPtr(15.0)}
	category := dto.MasterCategoryDto{ID: 1}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Camiseta", result[0].Name)
	assert.Equal(t, "Abrigo", result[1].Name)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterTypeService_GetTypesWithTempRangeAndCategory_MapsCategory(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockTypeRepository)

	// Prepare the mock data to return
	fakeData := []models.MasterType{
		{ID: 1, Name: "Camiseta", CategoryId: int64Ptr(3), MinTemp: nil, MaxTemp: nil},
	}
	mockRepo.On("GetTypesByCategory", mock.Anything).
		Return(fakeData, nil)

	service := services.NewMasterTypeService(mockRepo)
	weather := dto.WeatherDayDto{MaxTemp: floatPtr(25.0), MinTemp: floatPtr(15.0)}
	category := dto.MasterCategoryDto{ID: 3}

	// The function is executed
	result, err := service.GetTypesWithTempRangeAndCategory(weather, category)

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NotNil(t, result[0].Category)
	assert.Equal(t, int64(3), result[0].Category.ID)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}
