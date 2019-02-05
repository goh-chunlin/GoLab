package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"

	_ "github.com/lib/pq" // Create package-level variables and execute the init function of that package.
)

var db *sql.DB
var client appinsights.TelemetryClient

func main() {
	var err error
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	trace := appinsights.NewTraceTelemetry("Testing", appinsights.Information)
	trace.Timestamp = time.Now()

	client.Track(trace)

	// Initialize connection object.
	db, err = sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	checkError(err)

	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/addVideo", addVideo)
	mux.HandleFunc("/updateVideo", updateVideo)
	mux.HandleFunc("/deleteVideo", deleteVideo)

	err = http.ListenAndServe(getPort(), mux)
	checkError(err)

	client.TrackEvent("Client connected")
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}
