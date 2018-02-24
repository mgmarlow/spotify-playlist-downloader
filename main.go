package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mgmarlow/spotify-playlist-downloader/lib"
)

func main() {
	fmt.Println("Loading configuration file (./config.json)...")
	var config, err = lib.LoadConfig("./config.json")
	checkError("Error reading configuration file", err)

	fmt.Println("Generating Access Token...")
	token, err := lib.GetToken(config.ClientID + ":" + config.ClientSecret)
	checkError("Error generating token", err)
	fmt.Println("Access Token successfully created.")

	playlistURI, err := queryPlaylistURI()
	checkError("Error reading user input", err)

	trackItems, err := lib.GetAllTrackItems(playlistURI, token.AccessToken)
	checkError("Error fectching tracks from URI", err)

	err = lib.WriteToFile(trackItems)
	checkError("Unable to write tracks to a csv", err)
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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message+"\n", err)
	}
}
