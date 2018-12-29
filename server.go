package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	var err error

	// Initialize connection string.
	var connectionString = fmt.Sprintf(os.Getenv("CONNECTION_STRING"))

	// Initialize connection object.
	db, err = sql.Open("postgres", connectionString)
	checkError(err)

	mux := http.NewServeMux()

	staticFiles := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	mux.HandleFunc("/index", index)
	mux.HandleFunc("/addVideo", addVideo)
	mux.HandleFunc("/updateVideo", updateVideo)
	mux.HandleFunc("/deleteVideo", deleteVideo)

	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}

	server.ListenAndServe()
}
