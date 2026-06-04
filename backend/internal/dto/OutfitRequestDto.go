package dto

type OutfitRequestDto struct {
	City string `json:"city" binding:"required"`
	Days int64  `json:"days" binding:"required"`
}
