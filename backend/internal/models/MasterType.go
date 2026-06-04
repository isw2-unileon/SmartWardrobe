package models

type MasterType struct {
	ID         int64
	Name       string
	CategoryId *int64   `json:"category_id" gorm:"column:category_id"`
	MinTemp    *float64 `json:"min_temp" gorm:"column:min_temp"`
	MaxTemp    *float64 `json:"max_temp" gorm:"column:max_temp"`

	Category MasterCategory `gorm:"foreignKey:CategoryId"`
}

func (MasterType) TableName() string {
	return "master_types"
}
