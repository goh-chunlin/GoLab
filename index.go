package main

import (
	"mime"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "templates/index.html")

	writer.Header().Set("Content-Type", mime.TypeByExtension("html"))
}
