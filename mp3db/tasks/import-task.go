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

// ImportTask provides the ability to import all MP3 files
// from a directory and all subdirectories
type ImportTask struct {
	storage         storage.Storage
	fileAccess      files.FileAccess
	id3Reader       files.ID3Reader
	importLogWriter logger.ImportLogWriter
}

// NewImportTask - creates a new Instance of the ImportTask
func NewImportTask(storage storage.Storage, fileAccess files.FileAccess, id3Reader files.ID3Reader, importLogWriter logger.ImportLogWriter) *ImportTask {
	return &ImportTask{storage, fileAccess, id3Reader, importLogWriter}
}

// Execute - executes the task to import all MP3 files
func (task *ImportTask) Execute(rootDir string) (*model.ImportStats, error) {
	filesToImport, err := task.fileAccess.WalkDirectory(rootDir)
	if err != nil {
		return nil, err
	}

	failedFiles, err := task.startImport(filesToImport)
	if err != nil {
		return nil, err
	}

	importStats := model.NewImportStats(rootDir, filesToImport, failedFiles)

	task.importLogWriter.WriteLog(importStats)

	return importStats, nil
}

func (task *ImportTask) startImport(filesToImport *[]string) (*[]string, error) {
	totalCountOfFilesToImport := len(*filesToImport)

	fmt.Printf("Start with importing %v MP3 files...\n", totalCountOfFilesToImport)

	var failedFiles []string

	bar := progress.NewInt(totalCountOfFilesToImport)

	for index, fileToImport := range *filesToImport {
		song, err := task.id3Reader.Read(fileToImport)
		if err != nil {
			failedFiles = append(failedFiles, fileToImport)
			continue
		}

		err = task.storage.Insert(song)
		if err != nil {
			return nil, fmt.Errorf("Unexpected error while inserting into Songs table %v", err)
		}

		updateImportProgress(bar, index)
	}

	updateImportProgress(bar, totalCountOfFilesToImport)
	fmt.Println()

	return &failedFiles, nil
}

func updateImportProgress(bar *progress.Bar, index int) {
	bar.ValueInt(index)
	bar.WriteTo(os.Stdout)
}
