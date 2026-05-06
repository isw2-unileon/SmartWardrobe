package models

type MasterStyle struct {
	ID   int64
	Name string
}

func (MasterStyle) TableName() string {
	return "master_styles"
}
