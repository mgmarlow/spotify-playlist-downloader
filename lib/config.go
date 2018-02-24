package lib

import (
	"encoding/json"
	"os"
)

// Config contains clientID and secret from Spotify application
type Config struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// LoadConfig prepares basic auth application client ID and client secret
func LoadConfig(fileName string) (Config, error) {
	var config Config

	configFile, err := os.Open(fileName)
	if err != nil {
		return Config{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
