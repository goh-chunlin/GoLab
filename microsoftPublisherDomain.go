package main

import (
	"net/http"
	"strings"
)

func microsoftPublisherDomain(writer http.ResponseWriter, request *http.Request) {
	urlComponents := strings.Split(request.URL.Path, "/")

	http.ServeFile(writer, request, ".well-known/microsoft-identity-association.json")

	writer.Header().Set("Content-Type", "application/json")
}
