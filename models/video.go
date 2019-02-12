package models

import "database/sql"

// Video is a record of favourite video
type Video struct {
	ID             int    `json:"id"`
	Name           string `json:"videoTitle"`
	URL            string `json:"url"`
	YoutubeVideoID string `json:"youtubeVideoId"`
	CreatedBy      string `json:"createdBy"`
}

// GetVideo returns one single video record based on id
func GetVideo(id int) (video Video, err error) {
	video = Video{}

	err = db.QueryRow("SELECT id, name, url FROM videos WHERE id = $1;", id).Scan(&video.ID, &video.Name, &video.URL)
	video.YoutubeVideoID = video.URL[32:len(video.URL)]

	return
}

// GetAllVideos returns all video records
func GetAllVideos(userID string) (videos []Video, err error) {
	videos = []Video{}

	// Read data from table.
	sqlStatement := "SELECT * FROM videos ORDER BY id;"

	rows, err := db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		video := Video{}

		err := rows.Scan(&video.ID, &video.Name, &video.URL)
		video.YoutubeVideoID = video.URL[32:len(video.URL)]
		video.CreatedBy = userID

		if err == sql.ErrNoRows {

			err = nil

		} else if err == nil {

			videos = append(videos, video)

		}
	}

	return
}

// CreateVideo creates a new video record in the database
func (video *Video) CreateVideo() (err error) {
	sqlStatement := "INSERT INTO videos (name, url) VALUES ($1, $2);"

	_, err = db.Exec(sqlStatement, video.Name, video.URL)

	return
}

// UpdateVideo is to update an existing video record in the database
func (video *Video) UpdateVideo() (err error) {
	sqlStatement := "UPDATE videos SET name = $1 WHERE id = $2;"

	_, err = db.Exec(sqlStatement, video.Name, video.ID)

	return
}

// DeleteVideo is to delete an existing video record in the database
func (video *Video) DeleteVideo() (err error) {
	sqlStatement := "DELETE FROM videos WHERE id = $1;"

	_, err = db.Exec(sqlStatement, video.ID)

	return
}
