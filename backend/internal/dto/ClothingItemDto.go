package dto

type ClothingItemDto struct {
	ID       int64           `json:"id"`
	Type     *MasterTypeDto  `json:"type" binding:"required"`
	Color    *MasterColorDto `json:"color" binding:"required"`
	ImageUrl string          `json:"imageUrl" binding:"required"`
	Style    *MasterStyleDto `json:"style" binding:"required"`
}
