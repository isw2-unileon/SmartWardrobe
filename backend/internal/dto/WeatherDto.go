package dto

type WeatherDto struct {
	City    string  `json:"city"`
	MinTemp float64 `json:"minTemp"`
	MaxTemp float64 `json:"maxTemp"`
}
