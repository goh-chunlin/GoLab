package models

import (
	"database/sql"
	"time"
)

// Video is a record of favourite video
type Video struct {
	ID             int    `json:"id"`
	Name           string `json:"videoTitle"`
	URL            string `json:"url"`
	YoutubeVideoID string `json:"youtubeVideoId"`
}

// GetVideo returns one single video record based on id
func GetVideo(userID string, id int) (video Video, err error) {
	video = Video{}

	sqlStatement := "SELECT id, name, url FROM videos WHERE created_by = $1 AND id = $2;"

	err = db.QueryRow(sqlStatement, userID, id).Scan(&video.ID, &video.Name, &video.URL)
	video.YoutubeVideoID = video.URL[32:len(video.URL)]

	return
}

// GetAllVideos returns all video records
func GetAllVideos(userID string) (videos []Video, err error) {
	videos = []Video{}

	// Read data from table.
	sqlStatement := "SELECT id, name, url FROM videos WHERE created_by = $1 ORDER BY id;"

	rows, err := db.Query(sqlStatement, userID)

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
	sqlStatement := "INSERT INTO videos (name, url, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $3, $4);"

	_, err = db.Exec(sqlStatement, video.Name, video.URL, time.Now(), userID)

	return
}

// UpdateVideo is to update an existing video record in the database
func (video *Video) UpdateVideo(userID string) (err error) {
	sqlStatement := "UPDATE videos SET name = $1, updated_at = $2, updated_by = $3 WHERE id = $4;"

	_, err = db.Exec(sqlStatement, video.Name, time.Now(), userID, video.ID)

	return
}

// DeleteVideo is to delete an existing video record in the database
func (video *Video) DeleteVideo() (err error) {
	sqlStatement := "DELETE FROM videos WHERE id = $1;"

	_, err = db.Exec(sqlStatement, video.ID)

	return
}
