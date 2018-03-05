package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAllTrackItems returns all tracks from the provided playlist URI
func GetAllTrackItems(playlistURI string, accessToken string) ([]*Item, error) {
	var trackItems []*Item
	queryParams := strings.Split(playlistURI, ":")
	initialURI := "https://api.spotify.com/v1/users/" + queryParams[2] + "/playlists/" + queryParams[4] + "/tracks"

	initialTracks, err := getTracks(initialURI, accessToken)
	if err != nil {
		return nil, err
	}
	trackItems = append(trackItems, initialTracks.Items...)

	// Due to spotify's max return of 100 tracks, repeatedly request next page
	nextURI := initialTracks.Next
	for nextURI != "" {
		newTracks, err := getTracks(nextURI, accessToken)
		if err != nil {
			return nil, err
		}

		trackItems = append(trackItems, newTracks.Items...)
		nextURI = newTracks.Next
	}

	return trackItems, nil
}

func getTracks(uri string, accessToken string) (*Tracks, error) {
	resp, err := requestWithAuth(uri, accessToken)
	if err != nil {
		return nil, err
	}

	return unMarshallTracks(resp)
}

func unMarshallTracks(resp *http.Response) (*Tracks, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tracks Tracks
	err = json.Unmarshal(body, &tracks)
	return &tracks, nil
}

func requestWithAuth(path string, accessToken string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	return client.Do(req)
}
