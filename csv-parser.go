package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// WriteToFile writes the provided tracks to a csv, playlist.csv
func WriteToFile(trackItems []Item) error {
	// file, err := os.Create("playlist.csv")
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()

	w := csv.NewWriter(os.Stdout)

	for _, item := range trackItems {
		track := item.Track
		line := []string{
			track.Name,
			track.Album.Name,
			track.Artists[0].Name,
			item.AddedAt,
			strconv.Itoa(track.Popularity),
		}
		if err := w.Write(line); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}
