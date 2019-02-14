package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/goh-chunlin/GoLab/models"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/video/", handleVideoAPIRequests)
	writer = httptest.NewRecorder()
}

func TestHandleGetAllVideos(t *testing.T) {
	request, _ := http.NewRequest("GET", "post", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		// %v means the value in a default format
		//
		// Reference: https://golang.org/pkg/fmt/
		t.Errorf("Response code is %v", writer.Code)
	}

	var videos []models.Video
	json.Unmarshal(writer.Body.Bytes(), &videos)

	if len(videos) == 0 {
		t.Errorf("No video available")
	}
}
