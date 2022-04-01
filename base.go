package listrak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const HOST = "https://api.listrak.com"
const OAuth2 = "/OAuth2"

var clientID string
var clientSecret string
var bearerToken string

type baseRequest struct {
	params url.Values
	client http.Client
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func Initialize(id, secret string) (err error) {
	clientID = id
	clientSecret = secret
	bearerToken, err = getAuthToken()
	if err != nil {
		clientID = ""
		clientSecret = ""
	}
	return
}

// build client
//'headers' => ['Authorization' => 'Bearer ' . $accessToken],

// get auth token
func getAuthToken() (bearerToken string, err error) {
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
		err = fmt.Errorf("api.go getAuthToken 2")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("api.go getAuthToken 3")
		return
	}

	// Check if it has the 'access_token' info
	if !strings.Contains(string(body), "access_token") {
		err = fmt.Errorf("api.go getAuthToken 4")
		return
	}

	var tr TokenResponse

	err = json.Unmarshal(body, &tr)
	if err != nil {
		err = fmt.Errorf("api.go getAuthToken 5")
		return
	}

	bearerToken = tr.AccessToken

	if bearerToken == "" {
		err = fmt.Errorf("api.go getAuthToken 6")
		return
	}

	return bearerToken, err
}

// refresh if expired
func refreshToken() {}
