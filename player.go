package main

import (
	"mime"
	"net/http"
)

func player(writer http.ResponseWriter, request *http.Request) {
	user := profileFromSession(request)
	if user == nil {

		http.Redirect(writer, request, "/", http.StatusFound)

	} else {

		http.ServeFile(writer, request, "templates/player.html")

		writer.Header().Set("Content-Type", mime.TypeByExtension("html"))

	}
}
