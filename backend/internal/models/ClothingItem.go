package models

type ClothingItem struct {
	ID         int64
	TypeID     int64
	ColorID    int64
	StyleID    int64
	ImageURL   string
	UserID     string
	
	CategoryID int64
	Warmth     int64
	Layer      *int64
}