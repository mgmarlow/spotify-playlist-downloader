package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// TokenResponse provides the access token used to access the Spotify API
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// GetToken requests Spotify for an access token
func GetToken(basicAuth string) (TokenResponse, error) {
	resp, err := requestToken(basicAuth)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var token TokenResponse
	err = json.Unmarshal(body, &token)
	return token, err
}

func requestToken(basicAuth string) (*http.Response, error) {
	client := &http.Client{}
	form := url.Values{
		"grant_type": {"client_credentials"},
	}
	body := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}
