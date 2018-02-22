package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	var config, err = LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	client := config.ClientID + ":" + config.ClientSecret
	basicAuth := base64.StdEncoding.EncodeToString([]byte(client))
	token, err := GetToken(basicAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", token.AccessToken)
}
