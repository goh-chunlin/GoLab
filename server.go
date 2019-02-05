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

func main() {
	var err error

	// Initialize connection object.
	db, err = sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	checkError(err)

	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	mux.Handle("/", applicationInsightsLog(index))
	mux.HandleFunc("/addVideo", addVideo)
	mux.HandleFunc("/updateVideo", updateVideo)
	mux.HandleFunc("/deleteVideo", deleteVideo)

	err = http.ListenAndServe(getPort(), mux)
	checkError(err)
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}

func applicationInsightsLog(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		h(writer, request)
		duration := time.Now().Sub(startTime)

		client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

		trace := appinsights.NewRequestTelemetry(request.Method, request.URL.Path, duration, "200")
		trace.Timestamp = time.Now()

		client.Track(trace)
	})
}
