package services

import (
	"backend/internal/dto"
	"fmt"
	"math/rand"
)

// Interfaces de los servicios que necesita
type LocationServiceInterface interface {
	GetLocation(city string) (*dto.LocationDto, error)
}
type WeatherServiceInterface interface {
	GetWeather(location *dto.LocationDto, startDate string, endDate string) (*dto.WeatherDto, error)
}

type ClothingServiceInterface interface {
	GetClothingItem(filters dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error)
}

type MasterCategoriesServiceInterface interface {
	GetByName(name string) (*dto.MasterCategoryDto, error)
}

type MasterTypeServiceInterface interface {
	GetTypesWithTempRangeAndCategory(weather *dto.WeatherDto, category dto.MasterCategoryDto) ([]dto.MasterTypeDto, error)
}

type OutfitService struct {
	locationService         LocationServiceInterface
	weatherService          WeatherServiceInterface
	clothingService         ClothingServiceInterface
	masterCategoriesService MasterCategoriesServiceInterface
	masterTypeService       MasterTypeServiceInterface
}

func NewOutfitService(locationService LocationServiceInterface, weatherService WeatherServiceInterface, clothingService ClothingServiceInterface, categoryService MasterCategoriesServiceInterface, typeService MasterTypeServiceInterface) *OutfitService {
	return &OutfitService{
		locationService:         locationService,
		weatherService:          weatherService,
		clothingService:         clothingService,
		masterCategoriesService: categoryService,
		masterTypeService:       typeService,
	}
}

func (s *OutfitService) GenerateOutfit(req dto.OutfitRequestDto, user dto.UserDto) (*dto.OutfitResponseDto, error) {
	// Call the service to obtain the latitude and longitude of the city
	location, err := s.locationService.GetLocation(req.City)
	if err != nil {
		fmt.Printf("error getting location: %v\n", err)
		return nil, err
	}
	fmt.Printf("location: %v\n", location)

	// Call the WeatherService
	weather, err := s.weatherService.GetWeather(location, req.StartDate, req.EndDate)
	if err != nil {
		fmt.Printf("error getting weather: %v\n", err)
		return nil, err
	}
	fmt.Printf("weather: %+v\n", weather)

	// Generate the parts of the outfit
	upperwear := s.generateOutfitPart(weather, "upperwear", user)
	if upperwear == nil {
		return nil, fmt.Errorf("error search the upperwear part.")
	}

	bottomwear := s.generateOutfitPart(weather, "bottomwear", user)
	if bottomwear == nil {
		return nil, fmt.Errorf("error search the bottomwear part.")
	}

	footwear := s.generateOutfitPart(weather, "footwear", user)
	if footwear == nil {
		return nil, fmt.Errorf("error search the footwear part.")
	}

	outerwear := s.generateOutfitPart(weather, "outerwear", user)

	outfit := dto.OutfitDto{
		Upperwear:  upperwear,
		Bottomwear: bottomwear,
		Footwear:   footwear,
		Outerwear:  outerwear,
	}

	return &dto.OutfitResponseDto{
		Weather: *weather,
		Outfit:  outfit,
	}, nil
}

func (s *OutfitService) generateOutfitPart(weather *dto.WeatherDto, name string, user dto.UserDto) *dto.ClothingItemDto {
	// Call the MasterCategoryService with the name to obtain the id
	category, err := s.masterCategoriesService.GetByName(name)
	if err != nil {
		fmt.Printf("error getting category %s\n", name)
		return nil
	}

	// Call the MasterTypeService with the temperatures to obtain the clothing types can be wear
	types, err := s.masterTypeService.GetTypesWithTempRangeAndCategory(weather, *category)
	if err != nil {
		fmt.Printf("error getting type: %s\n", name)
		return nil
	}

	if len(types) == 0 {
		fmt.Printf("don't find types: %s\n", name)
		return nil
	}
	randomType := types[rand.Intn(len(types))]

	filters := dto.ClothingItemDto{
		Type: &randomType,
	}

	clothes, err := s.clothingService.GetClothingItem(filters, user)
	if err != nil {
		fmt.Printf("error getting clothes: %s\n", name)
		return nil
	}

	if len(clothes) == 0 {
		fmt.Printf("don't find clothes: %s\n", name)
		return nil
	}

	return &clothes[rand.Intn(len(clothes))]

}
