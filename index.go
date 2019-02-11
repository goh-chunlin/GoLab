package main

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/goh-chunlin/GoLab/models"
)

func index(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "templates/index.html")

	writer.Header().Set("Content-Type", mime.TypeByExtension("html"))
}

func indexWithJSON(writer http.ResponseWriter, request *http.Request) {
	video, err := models.GetVideo(8)
	checkError(err)

	output, err := json.MarshalIndent(&video, "", "\t\t")
	checkError(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}
