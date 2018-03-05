package lib

import (
	"encoding/csv"
	"os"
	"strconv"
)

// WriteToFile writes the provided tracks to a csv, playlist.csv
func WriteToFile(trackItems []*Item) error {
	file, err := os.Create("playlist.csv")
	if err != nil {
		return err
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range trackItems {
		track := item.Track
		line := []string{
			track.Name,
			track.Artists[0].Name,
			track.Album.Name,
			item.AddedAt,
			strconv.Itoa(track.Popularity),
		}

		if err := writer.Write(line); err != nil {
			return err
		}
	}

	return nil
}
