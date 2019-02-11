package main

import (
	"encoding/json"
	"models"
	"net/http"
	"path"
	"strconv"
	"util"
)

func handleVideoAPIRequests(writer http.ResponseWriter, request *http.Request) {
	var err error

	switch request.Method {
	case "GET":
		err = handleVideoAPIGet(writer, request)
	case "POST":
		err = handleVideoAPIPost(writer, request)
	case "PUT":
		err = handleVideoAPIPut(writer, request)
	case "DELETE":
		err = handleVideoAPIDelete(writer, request)
	}

	if err != nil {
		util.CheckError(err)
		return
	}
}

func handleVideoAPIGet(writer http.ResponseWriter, request *http.Request) (err error) {
	videoIDURL := path.Base(request.URL.Path)

	var output []byte

	if videoIDURL == "video" {
		videos, errIf := models.GetAllVideos()
		err = errIf
		util.CheckError(errIf)

		output, errIf = json.MarshalIndent(&videos, "", "\t")
		err = errIf
		util.CheckError(errIf)

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

		return
	}

	videoID, err := strconv.Atoi(videoIDURL)

	if err != nil {
		util.CheckError(err)
		return
	}

	video, err := models.GetVideo(videoID)
	util.CheckError(err)

	output, err = json.MarshalIndent(&video, "", "\t")
	util.CheckError(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}

func handleVideoAPIPost(writer http.ResponseWriter, request *http.Request) (err error) {
	length := request.ContentLength
	body := make([]byte, length)
	request.Body.Read(body)

	var video models.Video

	json.Unmarshal(body, &video)

	err = video.CreateVideo()
	util.CheckError(err)

	apiStatus := models.APIStatus{
		Status:  true,
		Message: "A video is successfully added to the database." + string(body[:length]),
	}
	output, err := json.MarshalIndent(&apiStatus, "", "\t")
	util.CheckError(err)

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}

func handleVideoAPIPut(writer http.ResponseWriter, request *http.Request) (err error) {
	videoIDURL := path.Base(request.URL.Path)

	videoID, err := strconv.Atoi(videoIDURL)

	if err != nil {
		util.CheckError(err)
		return
	}

	video, err := models.GetVideo(videoID)
	util.CheckError(err)

	length := request.ContentLength
	body := make([]byte, length)
	request.Body.Read(body)

	json.Unmarshal(body, &video)

	err = video.UpdateVideo()
	util.CheckError(err)

	apiStatus := models.APIStatus{
		Status:  true,
		Message: "A video record is successfully updated.",
	}
	output, err := json.MarshalIndent(&apiStatus, "", "\t")
	util.CheckError(err)

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}

func handleVideoAPIDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	videoIDURL := path.Base(request.URL.Path)

	videoID, err := strconv.Atoi(videoIDURL)

	if err != nil {
		util.CheckError(err)
		return
	}

	video, err := models.GetVideo(videoID)
	util.CheckError(err)

	err = video.DeleteVideo()
	util.CheckError(err)

	apiStatus := models.APIStatus{
		Status:  true,
		Message: "A video record is deleted.",
	}
	output, err := json.MarshalIndent(&apiStatus, "", "\t")
	util.CheckError(err)

	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}
