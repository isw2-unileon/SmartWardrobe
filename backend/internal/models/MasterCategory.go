package models

type MasterCategory struct {
	ID   int64
	Name string
}

func (MasterCategory) TableName() string {
	return "master_categories"
}
