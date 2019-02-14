package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
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
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video", handleVideoAPIRequests)
	writer = httptest.NewRecorder()
}

func TestHandleGetAllVideos(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/video", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		// %v means the value in a default format
		//
		// Reference: https://golang.org/pkg/fmt/
		t.Errorf("Response code is %v", writer.Code)
	}

	// var videos []models.Video
	// json.Unmarshal(writer.Body.Bytes(), &videos)

	// if len(videos) == 0 {
	// 	t.Errorf("No video available")
	// }
}
