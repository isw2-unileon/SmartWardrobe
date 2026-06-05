package dto

type OutfitRequestDto struct {
	City      string `json:"city" binding:"required"`
	Country   string `json:"country" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}
