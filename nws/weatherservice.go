package nws

import (
	"fmt"

	"github.com/benjamonnguyen/weather-service-takehome"
)

type WeatherService struct {
	apiClient ApiClient
}

func NewWeatherService() weather.WeatherService {
	return &WeatherService{
		apiClient: NewApiClient(),
	}
}

func (s *WeatherService) GetWeather(lat, long float32) (weather.WeatherResponse, error) {
	pointForecastsResp, err := s.apiClient.GetPointForecasts(lat, long)
	if err != nil {
		return weather.WeatherResponse{}, err
	}

	forecastResp, err := s.apiClient.GetForecast(pointForecastsResp.Properties.Forecast)
	if err != nil {
		return weather.WeatherResponse{}, err
	}

	periods := forecastResp.Properties.Periods
	if len(periods) == 0 {
		return weather.WeatherResponse{}, fmt.Errorf("no periods")
	}

	today := periods[0]
	return weather.WeatherResponse{
		TodaysForecast:              today.ShortForecast,
		TemperatureCharacterization: getTemperatureCharacterization(today.Temperature),
	}, nil
}

func getTemperatureCharacterization(temp int) string {
	if temp < 50 {
		return "cold"
	} else if temp < 80 {
		return "moderate"
	} else {
		return "hot"
	}
}
