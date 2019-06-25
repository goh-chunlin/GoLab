package main

import (
	"net/http"
)

func microsoftPublisherDomain(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"associatedApplications": [
		  {
			"applicationId": "332f4102-5c40-4f80-a70e-5023184125a1"
		  }
		]
	  }`)

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonData)
}
