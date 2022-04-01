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
const OAuth2 = "/OAuth2"

var clientID string
var clientSecret string
var tokenData AuthenticationResponse

type baseRequest struct {
	client  http.Client
	request *http.Request
}

//AuthenticationResponse from Listrak
type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	expiresAt   time.Time
}

//Initialize auth for bearerToken
func Initialize(id, secret string) (err error) {
	clientID = id
	clientSecret = secret
	err = setAuthToken()
	if err != nil {
		clientID = ""
		clientSecret = ""
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

	resp, err := http.Post(HOST+OAuth2+"/Token", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		err = fmt.Errorf("failed setting up auth request: %v", err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed reading body: %v", err.Error())
		return
	}

	var auth AuthenticationResponse
	err = json.Unmarshal(body, &auth)
	if err != nil {
		err = fmt.Errorf("failed reading the auth response: %v", err.Error())
		return
	}
	auth.expiresAt = time.Unix(auth.ExpiresIn, 0)
	tokenData = auth
	return
}

func getAccessToken() (token string, err error) {
	if tokenData.expiresAt.After(time.Now().Add(-5 * time.Minute)) {
		err = setAuthToken()
		if err != nil {
			err = fmt.Errorf("failed token refresh: %v", err.Error())
			return
		}
	}
	token = tokenData.AccessToken
	return
}

func (b baseRequest) setHeaders() {
	token, err := getAccessToken()
	if err != nil {
		return
	}
	b.request.Header.Set("Content-Type", "application/json")
	b.request.Header.Add("Authorization", "Bearer "+token)
}
