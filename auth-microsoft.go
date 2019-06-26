package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"

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
		Scopes:       []string{"https://graph.microsoft.com/User.Read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
		},
	}
	authCodeURL := config.AuthCodeURL(sessionID)

	http.Redirect(writer, request, authCodeURL, http.StatusFound)
}

// oauthCallbackWithMicrosoftHandler completes the OAuth flow, retreives the user's profile
// information from Microsoft Graph and stores it in a session.
func oauthCallbackWithMicrosoftHandler(writer http.ResponseWriter, request *http.Request) {
	oauthFlowSession, err := SessionStore.Get(request, request.FormValue("state"))
	if err != nil {
		util.CheckError(err)
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)
	// Validate this callback request came from the app.
	if !ok {
		util.CheckError(err)
	}

	code := request.FormValue("code")

	submission := &LoginSubmission{
		ClientID:     OAuthConfig.ClientID,
		ClientSecret: OAuthConfig.ClientSecret,
		GrandType:    "authorization_code",
		Code:         code,
		RedirectURI:  OAuthConfig.RedirectURL,
	}

	submissionJSON, err := json.MarshalIndent(&submission, "", "\t")
	if err != nil {
		util.CheckError(err)
	}

	req, err := http.NewRequest("POST", OAuthConfig.Endpoint.TokenURL, bytes.NewBuffer(submissionJSON))
	if err != nil {
		util.CheckError(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.CheckError(err)
	}
	defer resp.Body.Close()

	var tok *oauth2.Token
	body, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &tok); err != nil {
		util.CheckError(err)
	}

	clientAppInsights := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	trace := appinsights.NewTraceTelemetry("Body: "+string(body), appinsights.Information)
	trace.Timestamp = time.Now()

	clientAppInsights.Track(trace)
	// tok, err := OAuthConfig.Exchange(context.Background(), code)
	// if err != nil {
	// 	util.CheckError(err)
	// }

	session, err := SessionStore.New(request, defaultSessionID)
	if err != nil {
		util.CheckError(err)
	}

	ctx := context.Background()
	profile, err := fetchProfileFromMicrosoftGraph(ctx, tok)
	if err != nil {
		util.CheckError(err)
	}

	session.Values[oauthTokenSessionKey] = tok
	// Strip the profile to only the fields we need. Otherwise the struct is too big.
	session.Values[googleProfileSessionKey] = profile
	if err := session.Save(request, writer); err != nil {
		util.CheckError(err)
	}

	http.Redirect(writer, request, redirectURL, http.StatusFound)
}

// fetchProfile retrieves the Microsoft profile of the user associated with the
// provided OAuth token.
func fetchProfileFromMicrosoftGraph(ctx context.Context, tok *oauth2.Token) (*Profile, error) {
	client := oauth2.NewClient(ctx, OAuthConfig.TokenSource(ctx, tok))
	resp, err := client.Get("https://graph.microsoft.com/v1.0/users/me")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var microsoftGraphProfile map[string]interface{}

	data, _ := ioutil.ReadAll(resp.Body)

	clientAppInsights := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	trace := appinsights.NewTraceTelemetry(string(data), appinsights.Information)
	trace.Timestamp = time.Now()

	clientAppInsights.Track(trace)

	json.Unmarshal(data, &microsoftGraphProfile)

	return &Profile{
		ID:          microsoftGraphProfile["id"].(string),
		DisplayName: microsoftGraphProfile["displayName"].(string),
		ImageURL:    "",
	}, nil
}

type LoginSubmission struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrandType    string `json:"grand_type"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}
