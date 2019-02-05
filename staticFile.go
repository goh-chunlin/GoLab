package main

import (
	"net/http"
)

func staticFile(writer http.ResponseWriter, request *http.Request) {
	staticFiles := http.FileServer(http.Dir("public"))
	http.StripPrefix("/static/", staticFiles)
}
