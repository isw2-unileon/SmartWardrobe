package dto

type OutfitResponseDto struct {
	Weather WeatherDayDto `json:"weather" binding:"required"`
	Outfit  OutfitDto     `json:"outfit" binding:"required"`
}
