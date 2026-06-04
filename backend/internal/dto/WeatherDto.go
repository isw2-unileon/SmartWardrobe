package dto

type WeatherDto struct {
	Daily struct {
		Time    []string  `json:"time"`
		MaxTemp []float64 `json:"temperature_2m_max"`
		MinTemp []float64 `json:"temperature_2m_min"`
	} `json:"daily"`
}
