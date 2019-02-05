package main

import (
	"mime"
	"net/http"
	"strings"
)

func staticFile(writer http.ResponseWriter, request *http.Request) {
	staticFiles := http.FileServer(http.Dir("public"))
	http.StripPrefix("/static/", staticFiles)

	urlComponents := strings.Split(request.URL.Path, ".")
	fileExtension := urlComponents[len(urlComponents)-1]

	writer.Header().Set("Content-Type", mime.TypeByExtension(fileExtension))
}
