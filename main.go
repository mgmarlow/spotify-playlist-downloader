package main

import (
	"fmt"
	"log"
)

func main() {
	var config, err = LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	token, err := GetToken(config.ClientID + ":" + config.ClientSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", token.AccessToken)
}
