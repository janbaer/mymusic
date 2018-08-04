package files

import (
	"log"
	"os"

	"github.com/dhowden/tag"
	"github.com/janbaer/mp3db/model"
)

// MP3MetadataReader - defines the functions to read the MP3 metadata for the given file
type MP3MetadataReader interface {
	Read(filePath string) (*model.Song, error)
}

// RealMP3MetadataReader - implements the functions for MP3MetaDataReader
type RealMP3MetadataReader struct {
}

// Read - reads the mp3 metadata from the given file
func (reader RealMP3MetadataReader) Read(filePath string) (*model.Song, error) {
	mp3Metadata, err := readMp3File(filePath)
	if err != nil {
		return nil, err
	}

	return mapMetadataToSong(filePath, mp3Metadata), nil
}

func readMp3File(filename string) (tag.Metadata, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Could not read the following file", filename)
		return nil, err
	}

	mp3Metadata, err := tag.ReadFrom(f)
	if err != nil {
		// log.Println("Could not read any id3 tags from the following file", filename)
		return nil, err
	}

	return mp3Metadata, nil
}

func mapMetadataToSong(filePath string, mp3Metadata tag.Metadata) *model.Song {
	return &model.Song{
		FilePath: filePath,
		Artist:   mp3Metadata.Artist(),
		Title:    mp3Metadata.Title(),
		Album:    mp3Metadata.Album(),
		Genre:    mp3Metadata.Genre(),
	}
}
