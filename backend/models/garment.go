package models

type Garment struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Style    string `json:"style"`
	ImageURL string `json:"image_url"`
}