package files

import (
	"fmt"
	"log"

	"github.com/janbaer/mp3db/model"
	taglib "github.com/wtolson/go-taglib"
)

// ID3Reader - defines the interface for reading ID3 tags from a mp3 file
type ID3Reader interface {
	Read(filePath string) (*model.Song, error)
}

// ID3Writer - defines the interface for writing ID3 tags
type ID3Writer interface {
	Write(filePath string, song *model.Song) error
}

// ID3TagReadWriter - implements the functions for MP3MetaDataReader and MP3MetaDataWriter
type ID3TagReadWriter struct {
}

// Read - reads the mp3 metadata from the given file
func (reader ID3TagReadWriter) Read(filePath string) (*model.Song, error) {
	mp3File, err := readMp3File(filePath)
	if err != nil {
		return nil, err
	}

	if !verifyID3Tags(mp3File) {
		return nil, fmt.Errorf("No ID3 tags found for file %s", filePath)
	}

	return mapMetadataToSong(filePath, mp3File), nil
}

func (reader ID3TagReadWriter) Write(filePath string, song *model.Song) error {
	mp3File, err := readMp3File(filePath)

	if err != nil {
		mp3File.SetArtist(song.Artist)
		mp3File.SetAlbum(song.Album)
		mp3File.SetTitle(song.Title)
		mp3File.SetGenre(song.Genre)

		err = mp3File.Save()
	}

	return err
}

func readMp3File(filePath string) (*taglib.File, error) {
	mp3File, err := taglib.Read(filePath)
	if err != nil {
		log.Println("Could not read the following file", filePath)
		return nil, err
	}

	return mp3File, nil
}

func verifyID3Tags(mp3File *taglib.File) bool {
	if len(mp3File.Artist()) == 0 &&
		len(mp3File.Album()) == 0 &&
		len(mp3File.Title()) == 0 {
		return false
	}

	return true
}

func mapMetadataToSong(filePath string, mp3File *taglib.File) *model.Song {
	return &model.Song{
		FilePath: filePath,
		Artist:   mp3File.Artist(),
		Title:    mp3File.Title(),
		Album:    mp3File.Album(),
		Genre:    mp3File.Genre(),
	}
}
