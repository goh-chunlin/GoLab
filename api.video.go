package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"strconv"

	"github.com/goh-chunlin/GoLab/models"
	"github.com/goh-chunlin/GoLab/util"
)

func handleVideoAPIRequests(video models.IVideo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error

		user := profileFromSession(request)
		if user == nil {
			err = errors.New("sorry, you are not authorized")
			writer.WriteHeader(401)
			return
		}

		switch request.Method {
		case "GET":
			err = handleVideoAPIGet(writer, request, video, user)
		case "POST":
			err = handleVideoAPIPost(writer, request, video, user)
		case "PUT":
			err = handleVideoAPIPut(writer, request, video, user)
		case "DELETE":
			err = handleVideoAPIDelete(writer, request, video, user)
		}

		if err != nil {
			util.CheckError(err)
			return
		}
	}
}

func handleVideoAPIGet(writer http.ResponseWriter, request *http.Request, video models.IVideo, user *Profile) (err error) {

	videoIDURL := path.Base(request.URL.Path)

	var output []byte

	if videoIDURL == "video" {
		videos, errIf := video.GetAllVideos(user.ID)
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

	err = video.GetVideo(user.ID, videoID)
	util.CheckError(err)

	output, err = json.MarshalIndent(&video, "", "\t")
	util.CheckError(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)

	return
}

func handleVideoAPIPost(writer http.ResponseWriter, request *http.Request, video models.IVideo, user *Profile) (err error) {

	length := request.ContentLength
	body := make([]byte, length)
	request.Body.Read(body)

	json.Unmarshal(body, &video)

	err = video.CreateVideo(user.ID)

	if err != nil {

		util.CheckError(err)

		apiStatus := models.APIStatus{
			Status:  false,
			Message: err.Error(),
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(400)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	} else {

		apiStatus := models.APIStatus{
			Status:  true,
			Message: "A video is successfully added to the database.",
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	}

	return
}

func handleVideoAPIPut(writer http.ResponseWriter, request *http.Request, video models.IVideo, user *Profile) (err error) {

	videoIDURL := path.Base(request.URL.Path)

	videoID, err := strconv.Atoi(videoIDURL)

	if err != nil {
		util.CheckError(err)
		return
	}

	err = video.GetVideo(user.ID, videoID)
	if err != nil {

		util.CheckError(err)

		apiStatus := models.APIStatus{
			Status:  false,
			Message: err.Error(),
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(400)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	}

	length := request.ContentLength
	body := make([]byte, length)
	request.Body.Read(body)

	json.Unmarshal(body, &video)

	err = video.UpdateVideo(user.ID)
	if err != nil {

		util.CheckError(err)

		apiStatus := models.APIStatus{
			Status:  false,
			Message: err.Error(),
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(400)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	} else {

		apiStatus := models.APIStatus{
			Status:  true,
			Message: "A video record is successfully updated.",
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	}

	return
}

func handleVideoAPIDelete(writer http.ResponseWriter, request *http.Request, video models.IVideo, user *Profile) (err error) {

	videoIDURL := path.Base(request.URL.Path)

	videoID, err := strconv.Atoi(videoIDURL)

	if err != nil {
		util.CheckError(err)
		return
	}

	err = video.GetVideo(user.ID, videoID)
	if err != nil {

		util.CheckError(err)

		apiStatus := models.APIStatus{
			Status:  false,
			Message: err.Error(),
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(400)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)
	}

	err = video.DeleteVideo()
	if err != nil {

		util.CheckError(err)

		apiStatus := models.APIStatus{
			Status:  false,
			Message: err.Error(),
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(400)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	} else {

		apiStatus := models.APIStatus{
			Status:  true,
			Message: "A video record is deleted.",
		}
		output, err := json.MarshalIndent(&apiStatus, "", "\t")
		util.CheckError(err)

		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(output)

	}

	return
}
