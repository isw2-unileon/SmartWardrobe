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
type MockMasterCategoriesRepository struct {
	mock.Mock
}

func (m *MockMasterCategoriesRepository) GetByName(model models.MasterCategory) (models.MasterCategory, error) {
	args := m.Called(model)
	return args.Get(0).(models.MasterCategory), args.Error(1)
}

func TestMasterCategoriesService_GetByName_Success(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockMasterCategoriesRepository)

	// Prepare the mock data to return
	fakeCategory := models.MasterCategory{
		ID:   1,
		Name: "upperwear",
	}
	mockRepo.On("GetByName", mock.Anything).Return(fakeCategory, nil)

	svc := services.NewMasterCategoriesService(mockRepo)

	// The function is executed
	result, err := svc.GetByName("upperwear")

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "upperwear", result.Name)
	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterCategoriesService_GetByName_RepoError(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockMasterCategoriesRepository)

	// Prepare the mock error to return
	expectedError := errors.New("error fatal: database disconnected")
	mockRepo.On("GetByName", mock.Anything).Return(models.MasterCategory{}, expectedError)

	svc := services.NewMasterCategoriesService(mockRepo)

	// The function is executed
	result, err := svc.GetByName("upperwear")

	// The results are checked
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterCategoriesService_GetByName_MapsCorrectly(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockMasterCategoriesRepository)

	// Prepare the mock data to return
	fakeCategory := models.MasterCategory{
		ID:   3,
		Name: "footwear",
	}
	mockRepo.On("GetByName", mock.Anything).Return(fakeCategory, nil)

	svc := services.NewMasterCategoriesService(mockRepo)

	// The function is executed
	result, err := svc.GetByName("footwear")

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(3), result.ID)
	assert.Equal(t, "footwear", result.Name)
	assert.IsType(t, &dto.MasterCategoryDto{}, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}

func TestMasterCategoriesService_GetByName_PassesNameToRepo(t *testing.T) {
	// The fake repository is initialized
	mockRepo := new(MockMasterCategoriesRepository)

	// Prepare the mock data to return
	expectedModel := models.MasterCategory{Name: "bottomwear"}
	mockRepo.On("GetByName", expectedModel).Return(models.MasterCategory{
		ID:   2,
		Name: "bottomwear",
	}, nil)

	svc := services.NewMasterCategoriesService(mockRepo)

	// The function is executed
	result, err := svc.GetByName("bottomwear")

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verified that the service actually called the repository
	mockRepo.AssertExpectations(t)
}
