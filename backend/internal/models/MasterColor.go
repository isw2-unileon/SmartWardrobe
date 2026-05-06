package models

type MasterColor struct {
	ID   int64
	Name string
}

func (MasterColor) TableName() string {
	return "master_colors"
}
