package services_test

import (
	"backend/internal/dto"
	"backend/internal/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocks
type MockLocationService struct {
	mock.Mock
}

func (m *MockLocationService) GetLocation(city string, country string) (*dto.LocationDto, error) {
	args := m.Called(city, country)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*dto.LocationDto), args.Error(1)
}

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetWeather(location *dto.LocationDto, startDate string, endDate string) ([]dto.WeatherDayDto, error) {
	args := m.Called(location, startDate, endDate)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]dto.WeatherDayDto), args.Error(1)
}

type MockClothingService struct {
	mock.Mock
}

func (m *MockClothingService) GetClothingItem(filters dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error) {
	args := m.Called(filters, user)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]dto.ClothingItemDto), args.Error(1)
}

type MockMasterCategoriesService struct {
	mock.Mock
}

func (m *MockMasterCategoriesService) GetByName(name string) (*dto.MasterCategoryDto, error) {
	args := m.Called(name)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*dto.MasterCategoryDto), args.Error(1)
}

type MockMasterTypeService struct {
	mock.Mock
}

func (m *MockMasterTypeService) GetTypesWithTempRangeAndCategory(weather dto.WeatherDayDto, category dto.MasterCategoryDto) ([]dto.MasterTypeDto, error) {
	args := m.Called(weather, category)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]dto.MasterTypeDto), args.Error(1)
}

// Helpers
func buildMocks() (
	*MockLocationService,
	*MockWeatherService,
	*MockClothingService,
	*MockMasterCategoriesService,
	*MockMasterTypeService,
) {
	return new(MockLocationService),
		new(MockWeatherService),
		new(MockClothingService),
		new(MockMasterCategoriesService),
		new(MockMasterTypeService)
}

func defaultOutfitLocation() *dto.LocationDto {
	return &dto.LocationDto{
		Results: []struct {
			Name      string  `json:"name" binding:"required"`
			Country   string  `json:"country" binding:"required"`
			Latitude  float64 `json:"latitude" binding:"required"`
			Longitude float64 `json:"longitude" binding:"required"`
		}{
			{Name: "Madrid", Country: "Spain", Latitude: 40.4168, Longitude: -3.7038},
		},
	}
}

func defaultWeatherDays() []dto.WeatherDayDto {
	maxTemp := 25.0
	minTemp := 15.0
	return []dto.WeatherDayDto{
		{Date: "2026-06-04", MaxTemp: &maxTemp, MinTemp: &minTemp},
	}
}

func defaultClothingItem(name string) dto.ClothingItemDto {
	return dto.ClothingItemDto{
		Type:     &dto.MasterTypeDto{ID: 1, Name: name},
		ImageUrl: "https://imagen.com/" + name + ".jpg",
	}
}

// The mocks are set up
func setupHappyPath(
	locationSvc *MockLocationService,
	weatherSvc *MockWeatherService,
	clothingSvc *MockClothingService,
	categorySvc *MockMasterCategoriesService,
	typeSvc *MockMasterTypeService,
) {
	location := defaultOutfitLocation()
	weather := defaultWeatherDays()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(location, nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(weather, nil)

	categories := []string{"upperwear", "bottomwear", "footwear", "outerwear"}
	for i, cat := range categories {
		categorySvc.On("GetByName", cat).Return(&dto.MasterCategoryDto{ID: int64(i + 1), Name: cat}, nil)
		typeSvc.On("GetTypesWithTempRangeAndCategory", mock.Anything, mock.Anything).Return(
			[]dto.MasterTypeDto{{ID: int64(i + 1), Name: cat}}, nil,
		).Maybe()
		clothingSvc.On("GetClothingItem", mock.Anything, mock.Anything).Return(
			[]dto.ClothingItemDto{defaultClothingItem(cat)}, nil,
		).Maybe()
	}
}

func TestOutfitService_GenerateOutfit_Success(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()
	setupHappyPath(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{
		City:      "Madrid",
		Country:   "Spain",
		StartDate: "2026-06-04",
		EndDate:   "2026-06-04",
	}, dto.UserDto{ID: "user-uuid-123"})

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "2026-06-04", result[0].Weather.Date)
	assert.NotNil(t, result[0].Outfit.Upperwear)
	assert.NotNil(t, result[0].Outfit.Bottomwear)
	assert.NotNil(t, result[0].Outfit.Footwear)
}

func TestOutfitService_GenerateOutfit_MultipleDays(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	// Prepare the mock data to return
	maxTemp1, minTemp1 := 25.0, 15.0
	maxTemp2, minTemp2 := 20.0, 12.0
	weather := []dto.WeatherDayDto{
		{Date: "2026-06-04", MaxTemp: &maxTemp1, MinTemp: &minTemp1},
		{Date: "2026-06-05", MaxTemp: &maxTemp2, MinTemp: &minTemp2},
	}

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(weather, nil)

	categories := []string{"upperwear", "bottomwear", "footwear", "outerwear"}
	for i, cat := range categories {
		categorySvc.On("GetByName", cat).Return(&dto.MasterCategoryDto{ID: int64(i + 1), Name: cat}, nil)
	}
	typeSvc.On("GetTypesWithTempRangeAndCategory", mock.Anything, mock.Anything).
		Return([]dto.MasterTypeDto{{ID: 1, Name: "tipo"}}, nil)
	clothingSvc.On("GetClothingItem", mock.Anything, mock.Anything).
		Return([]dto.ClothingItemDto{defaultClothingItem("prenda")}, nil)

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{
		City: "Madrid", Country: "Spain",
		StartDate: "2026-06-04", EndDate: "2026-06-05",
	}, dto.UserDto{ID: "user-uuid-123"})

	// The results are checked
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "2026-06-04", result[0].Weather.Date)
	assert.Equal(t, "2026-06-05", result[1].Weather.Date)
}

func TestOutfitService_GenerateOutfit_LocationError(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).
		Return(nil, errors.New("city not found"))

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Inexistente"}, dto.UserDto{})

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "city not found")

	// Verified that the service called the correct functions
	locationSvc.AssertExpectations(t)
	weatherSvc.AssertNotCalled(t, "GetWeather")
}

func TestOutfitService_GenerateOutfit_WeatherError(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("weather API error"))

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Madrid"}, dto.UserDto{})

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verified that the service called the correct functions
	weatherSvc.AssertExpectations(t)
}

func TestOutfitService_GenerateOutfit_CategoryError(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(defaultWeatherDays(), nil)
	categorySvc.On("GetByName", mock.Anything).Return(nil, errors.New("category not found"))

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Madrid"}, dto.UserDto{})

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestOutfitService_GenerateOutfit_NoTypesFound(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(defaultWeatherDays(), nil)
	categorySvc.On("GetByName", mock.Anything).Return(&dto.MasterCategoryDto{ID: 1, Name: "upperwear"}, nil)
	typeSvc.On("GetTypesWithTempRangeAndCategory", mock.Anything, mock.Anything).
		Return([]dto.MasterTypeDto{}, nil)

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Madrid"}, dto.UserDto{})

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestOutfitService_GenerateOutfit_NoClothesFound(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(defaultWeatherDays(), nil)
	categorySvc.On("GetByName", mock.Anything).Return(&dto.MasterCategoryDto{ID: 1, Name: "upperwear"}, nil)
	typeSvc.On("GetTypesWithTempRangeAndCategory", mock.Anything, mock.Anything).
		Return([]dto.MasterTypeDto{{ID: 1, Name: "Camiseta"}}, nil)
	clothingSvc.On("GetClothingItem", mock.Anything, mock.Anything).
		Return([]dto.ClothingItemDto{}, nil)

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Madrid"}, dto.UserDto{})

	// The results are checked
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestOutfitService_GenerateOutfit_OuterwearIsOptional(t *testing.T) {
	locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc := buildMocks()

	locationSvc.On("GetLocation", mock.Anything, mock.Anything).Return(defaultOutfitLocation(), nil)
	weatherSvc.On("GetWeather", mock.Anything, mock.Anything, mock.Anything).Return(defaultWeatherDays(), nil)

	// Prepare the mock data to return
	for _, cat := range []string{"upperwear", "bottomwear", "footwear"} {
		categorySvc.On("GetByName", cat).Return(&dto.MasterCategoryDto{ID: 1, Name: cat}, nil)
	}

	categorySvc.On("GetByName", "outerwear").Return(nil, errors.New("not found"))

	typeSvc.On("GetTypesWithTempRangeAndCategory", mock.Anything, mock.Anything).
		Return([]dto.MasterTypeDto{{ID: 1, Name: "tipo"}}, nil)
	clothingSvc.On("GetClothingItem", mock.Anything, mock.Anything).
		Return([]dto.ClothingItemDto{defaultClothingItem("prenda")}, nil)

	svc := services.NewOutfitService(locationSvc, weatherSvc, clothingSvc, categorySvc, typeSvc)
	// The function is executed
	result, err := svc.GenerateOutfit(dto.OutfitRequestDto{City: "Madrid"}, dto.UserDto{})

	// The results are checked
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Nil(t, result[0].Outfit.Outerwear)
	assert.NotNil(t, result[0].Outfit.Upperwear)
	assert.NotNil(t, result[0].Outfit.Bottomwear)
	assert.NotNil(t, result[0].Outfit.Footwear)
}
