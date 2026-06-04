package services_test

import (
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
