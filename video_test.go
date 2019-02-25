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
	isTesting = true
}

func TestHandleGetAllVideos(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/api/video/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		// %v means the value in a default format
		//
		// Reference: https://golang.org/pkg/fmt/
		t.Errorf("Response code is %v", writer.Code)
	}

	var videos []models.Video
	json.Unmarshal(writer.Body.Bytes(), &videos)

	if len(videos) != 2 {
		t.Errorf("The list of videos is retrieved wrongly")
	}
}

func TestHandleGetOneVideoById(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/api/video/4", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var video models.Video
	json.Unmarshal(writer.Body.Bytes(), &video)

	if video.ID != 4 {
		t.Errorf("No video with ID = 4 available")
	}
}
