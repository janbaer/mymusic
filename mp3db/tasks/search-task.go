package tasks

import (
	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
)

// SearchTask - This task will search for the given search term in the storage
type SearchTask struct {
	storage storage.Storage
}

// NewSearchTask - creates a new instance of the SearchTask
func NewSearchTask(storage storage.Storage) *SearchTask {
	return &SearchTask{storage}
}

// Execute - Executes the taks and searches for the given search term
func (task *SearchTask) Execute(searchTerm string, searchOptions model.SearchOptions) (*[]model.Song, error) {
	searchQuery, values := storage.BuildSearchQuery(searchTerm, searchOptions)
	songs, err := task.storage.QuerySongs(searchQuery, values)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
