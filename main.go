package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/sujay1844/swiftie/swiftie"
)

//go:embed data/ts_lyrics.csv
var lyricsFile string

type Song struct {
	ID        string
	Name      string
	AlbumID   string
	AlbumName string
	AlbumPath string
	Lyrics    string
}

func initDB(reader io.Reader) ([]Song, error) {
	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	songs := make([]Song, 0, len(records))
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) != 6 {
			return nil, fmt.Errorf("record on line %d: expected 6 fields, got %d", i+1, len(record))
		}
		song := Song{
			ID:        record[0],
			Name:      record[1],
			AlbumID:   record[2],
			AlbumName: record[3],
			AlbumPath: record[4],
			Lyrics:    record[5],
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func main() {
	lyricsFileReader := strings.NewReader(lyricsFile)
	songs, err := initDB(lyricsFileReader)

	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}
	var song Song
	options := make([]huh.Option[Song], len(songs))
	for i, s := range songs {
		options[i] = huh.NewOption(s.Name, s)
	}
	huh.NewSelect[Song]().
		Title("Pick a country.").
		Options(options...).
		Value(&song).
		Filtering(true).
		WithHeight(10).
		Run()
}
