package models

type ClothingItem struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	TypeId   *int64 `json:"type_id" gorm:"column:type_id"`
	ColorId  *int64 `json:"color_id" gorm:"column:color_id"`
	ImageUrl string `json:"image_url" gorm:"column:image_url"`
	UserId   *int64 `json:"user_id" gorm:"column:user_id"`
	StyleId  *int64 `json:"style_id" gorm:"column:style_id"`
}

func (ClothingItem) TableName() string {
	return "clothing_items"
}
