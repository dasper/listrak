package listrak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const HOST = "https://api.listrak.com"
const AuthServer = "https://auth.listrak.com/OAuth2/Token"

var clientID string
var clientSecret string
var tokenData AuthenticationResponse

//AuthenticationResponse from Listrak
type AuthenticationResponse struct {
	expiresAt   time.Time
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

//Initialize auth for bearerToken
func Initialize(id, secret string) (err error) {
	clientID = id
	clientSecret = secret
	err = setAuthToken()
	if err != nil {
		clientID = ""
		clientSecret = ""
		return
	}
	return
}

// get auth token
func setAuthToken() (err error) {
	if clientID == "" || clientSecret == "" {
		err = ErrInvalidCredentials
		return
	}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)

	response, err := http.Post(AuthServer, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		err = fmt.Errorf("failed setting up auth request: %v", err.Error())
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed reading body: %v", err.Error())
		return
	}

	var auth AuthenticationResponse
	switch response.StatusCode {
	case 200:
		err = json.Unmarshal(body, &auth)
	case 404:
		err = ErrNotFound
	default:
		err = ErrUnhandledStatusCode
	}
	if err != nil {
		err = fmt.Errorf("failed reading the auth response: %v", err.Error())
		return
	}
	auth.expiresAt = time.Now().Add(time.Duration(auth.ExpiresIn) * time.Second)
	fmt.Println("Token expires at:", auth.expiresAt)
	tokenData = auth
	return
}

func getAccessToken() (token string, err error) {
	if tokenData.expiresAt.Before(time.Now().Add(-15 * time.Minute)) {
		err = setAuthToken()
		if err != nil {
			err = fmt.Errorf("failed token refresh: %v", err.Error())
			return
		}
	}
	token = tokenData.AccessToken
	return
}
