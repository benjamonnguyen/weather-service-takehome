package nws

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiClient interface {
	GetPointForecasts(lat, long float32) (PointForecastsResponse, error)
	GetForecast(endpoint string) (ForecastResponse, error)
}

type PointForecastsResponse struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type ForecastResponse struct {
	Properties struct {
		Periods []struct {
			ShortForecast string `json:"shortForecast"`
			Temperature   int    `json:"temperature"`
		} `json:"periods"`
	} `json:"properties"`
}

type apiClient struct {
	baseUrl string
}

func NewApiClient() ApiClient {
	return &apiClient{
		baseUrl: "https://api.weather.gov",
	}
}

func (c *apiClient) GetPointForecasts(lat, long float32) (PointForecastsResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/points/%.4f,%.4f", c.baseUrl, lat, long))
	if err != nil {
		return PointForecastsResponse{}, err
	}
	defer resp.Body.Close()

	// could write a util function to read and decode response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PointForecastsResponse{}, err
	}

	var res PointForecastsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return PointForecastsResponse{}, err
	}

	return res, nil
}

func (c *apiClient) GetForecast(endpoint string) (ForecastResponse, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return ForecastResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ForecastResponse{}, err
	}

	var res ForecastResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return ForecastResponse{}, err
	}

	return res, nil
}
