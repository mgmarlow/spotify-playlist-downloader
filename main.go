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
	checkError("Error reading from configuration file", err)

	fmt.Println("Generating Access Token...")
	token, err := GetToken(config.ClientID + ":" + config.ClientSecret)
	checkError("Error generating token", err)
	fmt.Println("Access Token successfully created.")

	playlistURI, err := queryPlaylistURI()
	checkError("Error reading user input", err)

	trackItems, err := GetAllTrackItems(playlistURI, token.AccessToken)
	checkError("Error fectching tracks from URI", err)

	err = WriteToFile(trackItems)
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
		log.Fatal(message, err)
	}
}
