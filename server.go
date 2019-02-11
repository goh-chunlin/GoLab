package main

import (
	"database/sql"
	"models"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Create package-level variables and execute the init function of that package.
)

var db *sql.DB

func main() {
	var err error

	// Initialize connection object.
	db, err = sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	checkError(err)

	models.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/static/", handleRequestWithLog(staticFile))

	mux.HandleFunc("/", handleRequestWithLog(index))
	mux.HandleFunc("/addVideo", handleRequestWithLog(addVideo))
	mux.HandleFunc("/updateVideo", handleRequestWithLog(updateVideo))
	mux.HandleFunc("/deleteVideo", handleRequestWithLog(deleteVideo))

	mux.HandleFunc("/api/video/", handleRequestWithLog(handleVideoAPIRequests))

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
