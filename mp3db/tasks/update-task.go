package tasks

import (
	"fmt"
	"os"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/logger"
	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/storage"
	progress "github.com/tj/go-progress"
)

// UpdateTask provides the ability to import all MP3 files
// from a directory and all subdirectories
type UpdateTask struct {
	storage           storage.Storage
	fileAccess        files.FileAccess
	mp3MetadataReader files.MP3MetadataReader
	updateLogWriter   logger.UpdateLogWriter
}

// NewImportTask - creates a new Instance of the UpdateTask
func NewUpdateTask(storage storage.Storage, fileAccess files.FileAccess, mp3MetadataReader files.MP3MetadataReader, updateLogWriter logger.UpdateLogWriter) *UpdateTask {
	return &UpdateTask{storage, fileAccess, mp3MetadataReader, updateLogWriter}
}

// Execute - executes the task to import all MP3 files
func (task *UpdateTask) Execute(rootDir string) (*model.UpdateStats, error) {
	filesToImport, err := task.fileAccess.WalkDirectory(rootDir)
	if err != nil {
		return nil, err
	}

	failedFiles, importFilesCount, updatedFilesCount, err := task.startUpdate(filesToImport)
	if err != nil {
		return nil, err
	}

	updateStats := model.NewUpdateStats(rootDir, filesToImport, importFilesCount, updatedFilesCount, failedFiles)

	task.updateLogWriter.WriteLog(updateStats)

	return updateStats, nil
}

func (task *UpdateTask) startUpdate(filesToImport *[]string) (*[]string, int, int, error) {
	totalCountOfFilesToImport := len(*filesToImport)

	fmt.Printf("Start with importing %v MP3 files...\n", totalCountOfFilesToImport)

	var failedFiles []string
	importFilesCount := 0
	updatedFilesCount := 0

	bar := progress.NewInt(totalCountOfFilesToImport)

	for index, fileToImport := range *filesToImport {
		song, err := task.mp3MetadataReader.Read(fileToImport)
		if err != nil {
			failedFiles = append(failedFiles, fileToImport)
			continue
		}

		songFromDb, _ := task.storage.QueryFilePath(fileToImport)
		if songFromDb == nil {
			importFilesCount++
			err = task.storage.Insert(song)
		} else if !songFromDb.TagsAreEqual(song) {
			updatedFilesCount++
			songFromDb.UpdateFrom(song)
			err = task.storage.Update(songFromDb)
		}

		if err != nil {
			return nil, 0, 0, fmt.Errorf("Unexpected error while inserting into Songs table %v", err)
		}

		updateUpdateProgress(bar, index)
	}

	updateImportProgress(bar, totalCountOfFilesToImport)
	fmt.Println()

	return &failedFiles, importFilesCount, updatedFilesCount, nil
}

func updateUpdateProgress(bar *progress.Bar, index int) {
	bar.ValueInt(index)
	bar.WriteTo(os.Stdout)
}
