package nws

import "testing"

var weatherSvc = &WeatherService{
	&mockApiClient{},
}

func TestGetWeather(t *testing.T) {
	resp, _ := weatherSvc.GetWeather(0, 0)
	if resp.TemperatureCharacterization != "cold" {
		t.Errorf("want cold, got %s", resp.TemperatureCharacterization)
	}
	want := "Cloudy, chance of meatballs"
	if resp.TodaysForecast != want {
		t.Errorf("want %s, got %s", want, resp.TodaysForecast)
	}
}

// not verifying args of mock calls
type mockApiClient struct{}

func (c *mockApiClient) GetPointForecasts(lat, long float32) (PointForecastsResponse, error) {
	return PointForecastsResponse{
		Properties: struct {
			Forecast string "json:\"forecast\""
		}{
			Forecast: "forecastendpoint",
		},
	}, nil
}

func (c *mockApiClient) GetForecast(endpoint string) (ForecastResponse, error) {
	return ForecastResponse{
		Properties: struct {
			Periods []struct {
				ShortForecast string "json:\"shortForecast\""
				Temperature   int    "json:\"temperature\""
			} "json:\"periods\""
		}{
			Periods: []struct {
				ShortForecast string "json:\"shortForecast\""
				Temperature   int    "json:\"temperature\""
			}{
				{
					ShortForecast: "Cloudy, chance of meatballs",
					Temperature:   -100,
				},
				{
					ShortForecast: "shortForecast",
					Temperature:   100,
				},
			},
		},
	}, nil
}
