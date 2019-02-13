package main

import (
	"mime"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	user := profileFromSession(request)
	if user != nil {

		http.ServeFile(writer, request, "templates/index.html")

		writer.Header().Set("Content-Type", mime.TypeByExtension("html"))

	} else {

		http.Redirect(writer, request, "/player", http.StatusFound)

	}
}
