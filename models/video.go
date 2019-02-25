package models

import (
	"database/sql"
	"errors"
	"time"
)

type IVideo interface {
	GetVideo(userID string, id int) (err error)
	GetAllVideos(userID string) (videos []Video, err error)
	CreateVideo(userID string) (err error)
	UpdateVideo(userID string) (err error)
	DeleteVideo() (err error)
}

// Video is a record of favourite video
type Video struct {
	Db             *sql.DB
	ID             int    `json:"id"`
	Name           string `json:"videoTitle"`
	URL            string `json:"url"`
	YoutubeVideoID string `json:"youtubeVideoId"`
}

// GetVideo returns one single video record based on id
func (video *Video) GetVideo(userID string, id int) (err error) {
	sqlStatement := "SELECT id, name, url FROM videos WHERE created_by = $1 AND id = $2;"

	err = video.Db.QueryRow(sqlStatement, userID, id).Scan(&video.ID, &video.Name, &video.URL)
	video.YoutubeVideoID = video.URL[32:len(video.URL)]

	return
}

// GetAllVideos returns all video records
func (video *Video) GetAllVideos(userID string) (videos []Video, err error) {
	videos = []Video{}

	// Read data from table.
	sqlStatement := "SELECT id, name, url FROM videos WHERE created_by = $1 ORDER BY id;"

	rows, err := video.Db.Query(sqlStatement, userID)

	defer rows.Close()

	for rows.Next() {
		video := Video{}

		err := rows.Scan(&video.ID, &video.Name, &video.URL)
		video.YoutubeVideoID = video.URL[32:len(video.URL)]

		if err == sql.ErrNoRows {

			err = nil

		} else if err == nil {

			videos = append(videos, video)

		}
	}

	return
}

// CreateVideo creates a new video record in the database
func (video *Video) CreateVideo(userID string) (err error) {
	if video.Name == "" {
		err = errors.New("the video name cannot be empty")

		return
	} else if video.URL == "" {
		err = errors.New("the video URL cannot be empty")

		return
	}
	sqlStatement := "INSERT INTO videos (name, url, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $3, $4);"

	_, err = video.Db.Exec(sqlStatement, video.Name, video.URL, time.Now(), userID)

	return
}

// UpdateVideo is to update an existing video record in the database
func (video *Video) UpdateVideo(userID string) (err error) {
	if video.Name == "" {
		err = errors.New("the video name cannot be empty")

		return
	}

	sqlStatement := "UPDATE videos SET name = $1, updated_at = $2, updated_by = $3 WHERE id = $4;"

	_, err = video.Db.Exec(sqlStatement, video.Name, time.Now(), userID, video.ID)

	return
}

// DeleteVideo is to delete an existing video record in the database
func (video *Video) DeleteVideo() (err error) {
	sqlStatement := "DELETE FROM videos WHERE id = $1;"

	_, err = video.Db.Exec(sqlStatement, video.ID)

	return
}
