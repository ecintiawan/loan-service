package main

import (
	"log"
	"time"

	app "github.com/ecintiawan/loan-service/internal/app/http"
)

func main() {
	// Load the desired time zone
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	// Set the default time zone
	time.Local = location

	// init http binary
	server := app.InitHttp()

	server.InitRoutes()
	server.ListenAndServe()
}
