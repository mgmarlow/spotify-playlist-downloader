package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type auth struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {
	var userAuth, err = loadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	client := userAuth.ClientID + ":" + userAuth.ClientSecret
	basicAuth := base64.StdEncoding.EncodeToString([]byte(client))
	token, err := getToken(basicAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", token.AccessToken)
}

func loadConfig(fileName string) (auth, error) {
	var userAuth auth

	configFile, err := os.Open(fileName)
	if err != nil {
		return auth{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&userAuth)
	if err != nil {
		return auth{}, err
	}

	return userAuth, nil
}

func getToken(basicAuth string) (tokenResponse, error) {
	resp, err := requestToken(basicAuth)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var token tokenResponse
	err = json.Unmarshal(body, &token)
	return token, err
}

// Client auth strategy
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
