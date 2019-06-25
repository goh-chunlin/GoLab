package main

import (
	"net/http"

	uuid "github.com/gofrs/uuid"
	"github.com/goh-chunlin/GoLab/util"
	"golang.org/x/oauth2"
)

func handleLoginWithMicrosoftRequest(writer http.ResponseWriter, request *http.Request) {
	sessionID := uuid.Must(uuid.NewV4()).String()

	oauthFlowSession, err := SessionStore.New(request, sessionID)
	util.CheckError(err)

	oauthFlowSession.Options.MaxAge = 3600 // 1 hour, i.e. 3,600 seconds

	oauthFlowSession.Values[oauthFlowRedirectKey] = "/player"

	if err := oauthFlowSession.Save(request, writer); err != nil {
		util.CheckError(err)
	}

	// Use the session ID for the "state" parameter.
	// This protects against CSRF (cross-site request forgery).
	// See https://godoc.org/golang.org/x/oauth2#Config.AuthCodeURL for more detail.
	config := &oauth2.Config{
		ClientID:     "332f4102-5c40-4f80-a70e-5023184125a1",
		ClientSecret: "/-/3w3hNvu@sFOnuH2DqJYLFAnAh1C2n",
		RedirectURL:  "https://golab002.azurewebsites.net/auth",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
		},
	}
	authCodeURL := config.AuthCodeURL(sessionID, oauth2.ApprovalForce,
		oauth2.AccessTypeOnline)

	http.Redirect(writer, request, authCodeURL, http.StatusFound)
}
