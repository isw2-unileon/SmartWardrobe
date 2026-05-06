package models

type WeatherRule struct {
	ID             int64   `db:"id"`
	MinTemp        float64 `db:"min_temp"`
	MaxTemp        float64 `db:"max_temp"`
	RequiredWarmth int64   `db:"required_warmth"`
	MaxUpperLayers int64   `db:"max_upper_layers"`
}