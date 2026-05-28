package dto

type ClothingItemDto struct {
	ID       int64  `json:"id"`
	TypeId   *int64 `json:"typeId" binding:"required"`
	ColorId  *int64 `json:"colorId" binding:"required"`
	ImageUrl string `json:"imageUrl" binding:"required"`
	StyleId  *int64 `json:"styleId" binding:"required"`
}
