package dto

type MasterColorDto struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
