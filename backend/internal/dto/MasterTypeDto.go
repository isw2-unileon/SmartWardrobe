package dto

type MasterTypeDto struct {
	ID       int64              `json:"id" binding:"required"`
	Name     string             `json:"name" binding:"required"`
	Category *MasterCategoryDto `json:"category"`
}
