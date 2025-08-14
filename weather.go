package weather

type WeatherService interface {
	GetWeather(lat, long float32) (WeatherResponse, error)
}

type WeatherResponse struct {
	TodaysForecast              string `json:"todaysForecast"`
	TemperatureCharacterization string `json:"temperatureCharacterization"` // can be enum
}
