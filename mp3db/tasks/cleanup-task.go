package tasks

import (
	"fmt"
	"os"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
	progress "github.com/tj/go-progress"
)

// CleanupTask - task for cleaning all songs their files no longer exists
type CleanupTask struct {
	storage    storage.Storage
	fileAccess files.FileAccess
}

// NewCleanupTask - creates a new instance of the CleanupTask
func NewCleanupTask(storage storage.Storage, fileAccess files.FileAccess) *CleanupTask {
	return &CleanupTask{storage, fileAccess}
}

// Execute - executes the task to cleanup the database from orphaned songs
func (task CleanupTask) Execute() (*model.CleanupStats, error) {
	songs, err := task.storage.QueryAll()
	if err != nil {
		return nil, err
	}

	totalCountOfSongs := len(*songs)
	bar := progress.NewInt(totalCountOfSongs)

	var deletedSongs []model.Song

	for index, song := range *songs {
		if !task.fileAccess.ExistsFile(song.FilePath) {
			deletedSongs = append(deletedSongs, song)
			task.storage.Delete(&song)
		}
		updateCleanupProgress(bar, index)
	}

	updateCleanupProgress(bar, totalCountOfSongs)
	fmt.Println()

	totalCountOfDeletedSongs := len(deletedSongs)
	stats := model.NewCleanupStats(totalCountOfSongs, totalCountOfDeletedSongs, &deletedSongs)

	return stats, nil
}

func updateCleanupProgress(bar *progress.Bar, index int) {
	bar.ValueInt(index)
	bar.WriteTo(os.Stdout)
}
