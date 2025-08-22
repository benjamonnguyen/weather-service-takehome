package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/benjamonnguyen/weather-service-takehome/nws"
)

func main() {
	// routing
	weatherCtrl := &weatherController{
		weatherSvc: nws.NewWeatherService(),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Weather Service"))
	})
	mux.HandleFunc("/forecast/{lat}/{long}", weatherCtrl.GetForecast)

	// middleware
	handler := authMiddleware(mux)
	handler = loggingMiddleware(handler)

	// server
	errCh := make(chan error)
	addr := ":8080"
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		errCh <- srv.ListenAndServe()
	}()
	log.Printf("listening at %s", addr)

	err := <-errCh
	log.Fatal(err)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Printf("Completed in %v\n", time.Since(start))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// DO AUTH STUFF
		next.ServeHTTP(w, r)
	})
}
