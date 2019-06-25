package main

import (
	"net/http"
)

func microsoftPublisherDomain(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	http.ServeFile(writer, request, ".well-known/microsoft-identity-association.json")
}
