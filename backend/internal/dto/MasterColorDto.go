package dto

type MasterColorDto struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
