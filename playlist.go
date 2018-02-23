package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Artist struct {
	Href string
	ID   string
	Name string
	Type string
	URI  string
}

type Album struct {
	AlbumType string
	Artists   []Artist
}

type Track struct {
	Album
}

type TrackItem struct {
	AddedAt string
	Track
}

type Tracks struct {
	Href     string
	Items    []TrackItem
	Next     string
	Offset   int
	Previous int
	Total    int
}

type Playlist struct {
	Collaborative bool
	Description   string
	Href          string
	ID            string
	Name          string
	Tracks        Tracks
}

func GetPlaylist(playlistURI string, accessToken string) (Playlist, error) {
	resp, err := requestPlaylist(playlistURI, accessToken)
	if err != nil {
		return Playlist{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var playlist Playlist
	err = json.Unmarshal(body, &playlist)
	return playlist, err
}

func requestPlaylist(playlistURI string, accessToken string) (*http.Response, error) {
	queryParams := strings.Split(playlistURI, ":")
	uri := "https://api.spotify.com/v1/users/" + queryParams[2] + "/playlists/" + queryParams[4]

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	return client.Do(req)
}
