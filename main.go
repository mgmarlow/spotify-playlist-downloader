package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type auth struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func main() {
	var userAuth, err = loadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	client := userAuth.ClientID + ":" + userAuth.ClientSecret
	basicAuth := base64.StdEncoding.EncodeToString([]byte(client))
	fmt.Printf("%s", basicAuth)
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
