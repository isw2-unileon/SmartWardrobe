package dto

type WeatherDayDto struct {
	Date    string   `json:"date"`
	MaxTemp *float64 `json:"maxTemp"`
	MinTemp *float64 `json:"minTemp"`
}
