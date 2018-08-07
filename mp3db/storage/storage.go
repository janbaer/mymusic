package storage

import "github.com/janbaer/mp3db/model"

// Storage - defines the interface for the access the storage where the infos about the mp3 files
// are stored
type Storage interface {
	QueryAll() (*[]model.Song, error)
	QuerySongs(query string, queryValues []string) (*[]model.Song, error)
	QueryByID(id int) (*model.Song, error)
	QueryFilePath(filePath string) (*model.Song, error)
	Insert(song *model.Song) error
	Update(song *model.Song) error
	Delete(song *model.Song) error
}
