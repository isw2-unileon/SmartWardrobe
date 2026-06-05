package dto

type LocationDto struct {
	Results []struct {
		Name      string  `json:"name" binding:"required"`
		Country   string  `json:"country" binding:"required"`
		Latitude  float64 `json:"latitude" binding:"required"`
		Longitude float64 `json:"longitude" binding:"required"`
	} `json:"results"`
}
