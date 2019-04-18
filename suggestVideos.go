package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/goh-chunlin/GoLab/models"
)

func suggestVideos(video models.IVideo) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		user := profileFromSession(request)
		if user == nil {
			writer.WriteHeader(401)
			return
		}

		go retrieveSuggestedVideos("anime", video, user)
	}
}

func retrieveSuggestedVideos(area string, video models.IVideo, user *Profile) {
	// YouTube API: https://developers.google.com/youtube/v3/docs/search/list
	// Golang: 	https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
	//			https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c

	youtubeDataURL := "https://www.googleapis.com/youtube/v3/search?"
	youtubeDataURL += "part=snippet&"
	youtubeDataURL += "maxResults=25&"
	youtubeDataURL += "topicId=%2Fm%2F04rlf&"
	youtubeDataURL += "type=video&"
	youtubeDataURL += "videoCategoryId=10&"
	youtubeDataURL += "videoDuration=medium&"
	youtubeDataURL += "q=" + area + "&"
	youtubeDataURL += "key=" + os.Getenv("YOUTUBE_API_KEY")

	request, err := http.NewRequest("GET", youtubeDataURL, nil)

	checkError(err)

	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	checkError(err)

	defer response.Body.Close()

	var youtubeResponse YoutubeResponse
	err = json.NewDecoder(response.Body).Decode(&youtubeResponse)
	checkError(err)

	for _, youtubeItem := range youtubeResponse.Items {
		video.CreateVideoWithNameAndURL(user.ID, youtubeItem.Snippet.Title, "https://www.youtube.com/watch?v="+youtubeItem.ItemID.VideoID)
	}
}

// YoutubeResponse is a data structure
type YoutubeResponse struct {
	Items []YoutubeItem `json:"items"`
}

// YoutubeItem is a data structure
type YoutubeItem struct {
	ItemID  YoutubeItemID      `json:"id"`
	Snippet YoutubeItemSnippet `json:"snippet"`
}

// YoutubeItemID is a data structure
type YoutubeItemID struct {
	VideoID string `json:"videoId"`
}

// YoutubeItemSnippet is a data structure
type YoutubeItemSnippet struct {
	Title string `json:"title"`
}
