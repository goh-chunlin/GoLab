package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// FakeVideo is a record of favourite video for unit test
type FakeVideo struct {
	ID             int    `json:"id"`
	Name           string `json:"videoTitle"`
	URL            string `json:"url"`
	YoutubeVideoID string `json:"youtubeVideoId"`
	CreatedBy      string `json:"createdBy"`
}

// GetVideo returns one single video record based on id
func (video *FakeVideo) GetVideo(userID string, id int) (err error) {
	jsonFile, err := os.Open("testdata/fake_videos.json")

	if err != nil {
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return
	}

	var fakeVideos []FakeVideo
	json.Unmarshal(jsonData, &fakeVideos)

	for _, fakeVideo := range fakeVideos {
		if fakeVideo.ID == id && fakeVideo.CreatedBy == userID {

			video.ID = fakeVideo.ID
			video.Name = fakeVideo.Name
			video.URL = fakeVideo.URL
			video.YoutubeVideoID = fakeVideo.YoutubeVideoID

			return
		}
	}

	err = errors.New("no corresponding video found")

	return
}

// GetAllVideos returns all video records
func (video *FakeVideo) GetAllVideos(userID string) (videos []Video, err error) {
	jsonFile, err := os.Open("testdata/fake_videos.json")

	if err != nil {
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return
	}

	var fakeVideos []FakeVideo
	json.Unmarshal(jsonData, &fakeVideos)

	for _, fakeVideo := range fakeVideos {
		if fakeVideo.CreatedBy == userID {
			video := Video{
				ID:             fakeVideo.ID,
				Name:           fakeVideo.Name,
				URL:            fakeVideo.URL,
				YoutubeVideoID: fakeVideo.YoutubeVideoID,
			}

			videos = append(videos, video)
		}
	}

	return
}

// CreateVideo creates a new video record in the database
func (video *FakeVideo) CreateVideo(userID string) (err error) {
	if video.Name == "" {
		err = errors.New("the video name cannot be empty")

		return
	} else if video.URL == "" {
		err = errors.New("the video URL cannot be empty")

		return
	}

	err = nil

	return
}

// UpdateVideo is to update an existing video record in the database
func (video *FakeVideo) UpdateVideo(userID string) (err error) {
	if video.Name == "" {
		err = errors.New("the video name cannot be empty")

		return
	}

	err = nil

	return
}

// DeleteVideo is to delete an existing video record in the database
func (video *FakeVideo) DeleteVideo() (err error) {

	err = nil

	return
}
