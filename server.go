package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/lib/pq"
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

	mux.HandleFunc("/", index)
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
