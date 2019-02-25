package main

import (
	"bytes"
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

func TestHandleCreateNewVideo(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{
		"videoTitle": "Lanota - Dreams Go On",
		"url": "https://www.youtube.com/watch?v=55H2rt1zy2g"
		}`)

	request, _ := http.NewRequest("POST", "/api/video/", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleCreateNewVideoWithoutVideoTitle(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{
		"url": "https://www.youtube.com/watch?v=55H2rt1zy2g"
		}`)

	request, _ := http.NewRequest("POST", "/api/video/", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleCreateNewVideoWithoutUrl(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{
		"videoTitle": "Lanota - Dreams Go On"
		}`)

	request, _ := http.NewRequest("POST", "/api/video/", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleCreateNewVideoWithoutVideoTitleAndUrl(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{}`)

	request, _ := http.NewRequest("POST", "/api/video/", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleUpdateVideo(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{
		"videoTitle": "Lanota - Dreams Go On"
	}`)

	request, _ := http.NewRequest("PUT", "/api/video/4", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleUpdateVideoWithoutTitle(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{}`)

	request, _ := http.NewRequest("PUT", "/api/video/4", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleUpdateInvalidVideo(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	var data = []byte(`{
		"videoTitle": "Lanota - Dreams Go On"
	}`)

	request, _ := http.NewRequest("PUT", "/api/video/2", bytes.NewBuffer(data))
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleDeleteVideo(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", "/api/video/4", nil)
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleDeleteInvalidVideo(t *testing.T) {
	mux = http.NewServeMux()
	mux.HandleFunc("/api/video/", handleVideoAPIRequests(&models.FakeVideo{}))
	writer = httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", "/api/video/2", nil)
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
