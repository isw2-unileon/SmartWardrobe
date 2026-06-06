package services

import (
	"backend/internal/dto"
	"fmt"
	"math/rand"
)

// Interfaces of the services that needs
type LocationServiceInterface interface {
	GetLocation(city string, country string) (*dto.LocationDto, error)
}
type WeatherServiceInterface interface {
	GetWeather(location *dto.LocationDto, startDate string, endDate string) ([]dto.WeatherDayDto, error)
}

type ClothingServiceInterface interface {
	GetClothingItem(filters dto.ClothingItemDto, user dto.UserDto) ([]dto.ClothingItemDto, error)
}

type MasterCategoriesServiceInterface interface {
	GetByName(name string) (*dto.MasterCategoryDto, error)
}

type MasterTypeServiceInterface interface {
	GetTypesWithTempRangeAndCategory(weather dto.WeatherDayDto, category dto.MasterCategoryDto) ([]dto.MasterTypeDto, error)
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

func (s *OutfitService) GenerateOutfit(req dto.OutfitRequestDto, user dto.UserDto) ([]dto.OutfitResponseDto, error) {
	// Call the service to obtain the latitude and longitude of the city
	location, err := s.locationService.GetLocation(req.City, req.Country)
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

	// Generate the parts of the outfit
	var outfits []dto.OutfitResponseDto
	for _, day := range weather {
		upperwear := s.generateOutfitPart(day, "upperwear", user)
		if upperwear == nil {
			return nil, fmt.Errorf("error searching the upperwear part")
		}

		bottomwear := s.generateOutfitPart(day, "bottomwear", user)
		if bottomwear == nil {
			return nil, fmt.Errorf("error searching the bottomwear part")
		}

		footwear := s.generateOutfitPart(day, "footwear", user)
		if footwear == nil {
			return nil, fmt.Errorf("error searching the footwear part")
		}

		outerwear := s.generateOutfitPart(day, "outerwear", user)

		outfits = append(outfits, dto.OutfitResponseDto{
			Weather: day,
			Outfit: dto.OutfitDto{
				Upperwear:  upperwear,
				Bottomwear: bottomwear,
				Footwear:   footwear,
				Outerwear:  outerwear,
			},
		})

	}

	return outfits, nil
}

func (s *OutfitService) generateOutfitPart(weather dto.WeatherDayDto, name string, user dto.UserDto) *dto.ClothingItemDto {
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

	var clothes []dto.ClothingItemDto
	for _, t := range types {
		filter := dto.ClothingItemDto{
			Type: &t,
		}
		clothingItem, err := s.clothingService.GetClothingItem(filter, user)
		if err != nil {
			fmt.Printf("error getting clothes: %s\n", name)
			return nil
		}
		clothes = append(clothes, clothingItem...)
	}

	if len(clothes) == 0 {
		fmt.Printf("don't find clothes: %s\n", name)
		return nil
	}

	return &clothes[rand.Intn(len(clothes))]

}
