package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/benjamonnguyen/weather-service-takehome"
)

type weatherController struct {
	weatherSvc weather.WeatherService
}

func (c *weatherController) GetForecast(w http.ResponseWriter, r *http.Request) {
	latVal := r.PathValue("lat")
	longVal := r.PathValue("long")
	lat, err := strconv.ParseFloat(latVal, 32)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest) // not returning error messages for now
	}
	long, err := strconv.ParseFloat(longVal, 32)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	resp, err := c.weatherSvc.GetWeather(float32(lat), float32(long))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(data)
}
