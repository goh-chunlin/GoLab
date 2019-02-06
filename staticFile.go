package main

import (
	"mime"
	"net/http"
	"strings"
)

func staticFile(writer http.ResponseWriter, request *http.Request) {
	urlComponents := strings.Split(request.URL.Path, "/")

	http.ServeFile(writer, request, "public/"+urlComponents[len(urlComponents)-1])

	fileComponents := strings.Split(urlComponents[len(urlComponents)-1], ".")
	fileExtension := fileComponents[len(fileComponents)-1]

	writer.Header().Set("Content-Type", mime.TypeByExtension(fileExtension))
}
