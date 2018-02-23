package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
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
func GetToken(client string) (TokenResponse, error) {
	base64Client := base64.StdEncoding.EncodeToString([]byte(client))
	resp, err := requestToken(base64Client)
	if err != nil {
		return TokenResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var token TokenResponse
	err = json.Unmarshal(body, &token)
	return token, err
}

func requestToken(base64Client string) (*http.Response, error) {
	client := &http.Client{}
	form := url.Values{
		"grant_type": {"client_credentials"},
	}
	body := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+base64Client)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(req)
}
