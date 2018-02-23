package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Loading configuration file (./config.json)...")
	var config, err = LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating Access Token...")
	token, err := GetToken(config.ClientID + ":" + config.ClientSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Access Token successfully created.")

	playlistURI, err := queryPlaylistURI()
	if err != nil {
		log.Fatal(err)
	}

	tracks, err := GetTracks(playlistURI, token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(tracks.Items); i++ {
		fmt.Printf("%s\n", tracks.Items[i].Track.Name)
	}
}

func queryPlaylistURI() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Playlist URI: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
