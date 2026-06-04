package dto

type OutfitResponseDto struct {
	Weather WeatherDto `json:"weather" binding:"required"`
	Outfit  OutfitDto  `json:"outfit" binding:"required"`
}
