package dto

type OutfitDto struct {
	Upperwear  *ClothingItemDto `json:"upperwear" binding:"required"`
	Bottomwear *ClothingItemDto `json:"bottomwear" binding:"required"`
	Footwear   *ClothingItemDto `json:"footwear" binding:"required"`
	Outerwear  *ClothingItemDto `json:"outerwear"`
}
