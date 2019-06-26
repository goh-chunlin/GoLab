package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/goh-chunlin/GoLab/models"

	_ "github.com/lib/pq" // Create package-level variables and execute the init function of that package.
)

func main() {
	// Initialize connection object.
	db, err := sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	checkError(err)

	models.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/static/", handleRequestWithLog(staticFile))

	mux.HandleFunc("/", handleRequestWithLog(index))
	mux.HandleFunc("/suggestVideos", handleRequestWithLog(suggestVideos(&models.Video{Db: db})))
	mux.HandleFunc("/player", handleRequestWithLog(player))
	mux.HandleFunc("/login", handleRequestWithLog(handleLoginRequest))
	mux.HandleFunc("/logout", handleRequestWithLog(handleLogoutRequest))
	mux.HandleFunc("/oauth2callback", handleRequestWithLog(oauthCallbackHandler))

	mux.HandleFunc("/login-with-microsoft", handleRequestWithLog(handleLoginWithMicrosoftRequest))
	mux.HandleFunc("/oauth2callback-with-microsoft", handleRequestWithLog(oauthCallbackWithMicrosoftHandler))
	mux.HandleFunc("/.well-known/microsoft-identity-association.json", handleRequestWithLog(microsoftPublisherDomain))

	mux.HandleFunc("/api/video/", handleRequestWithLog(handleVideoAPIRequests(&models.Video{Db: db})))

	err = http.ListenAndServe(getPort(), mux)
	checkError(err)
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}
