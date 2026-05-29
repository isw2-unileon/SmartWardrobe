package models

type ClothingItem struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	TypeId   *int64 `gorm:"column:type_id"`
	ColorId  *int64 `gorm:"column:color_id"`
	ImageUrl string `gorm:"column:image_url"`
	UserId   string `gorm:"column:user_id"`
	StyleId  *int64 `gorm:"column:style_id"`

	Type  MasterType  `gorm:"foreignKey:TypeId"`
	Color MasterColor `gorm:"foreignKey:ColorId"`
	Style MasterStyle `gorm:"foreignKey:StyleId"`
}

func (ClothingItem) TableName() string {
	return "clothing_items"
}
