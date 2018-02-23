package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	playlist, err := GetPlaylist(playlistURI, token.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Tracks for Playlist " + playlist.Name)

	for i := 0; i < len(playlist.Tracks.Items); i++ {
		fmt.Printf("%s", playlist.Tracks.Items[i].Track.Name)
	}
}

func queryPlaylistURI() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Playlist URI: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return text, nil
}
