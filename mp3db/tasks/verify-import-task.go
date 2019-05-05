package tasks

import (
	"fmt"
	"os"

	"github.com/janbaer/mp3db/files"
	"github.com/janbaer/mp3db/logger"
	"github.com/janbaer/mp3db/model"
	"github.com/janbaer/mp3db/web"
	progress "github.com/tj/go-progress"
)

// VerifyImportTask provides the ability to verify all MP3 files
// from the given directory before importing the prevent duplicates
type VerifyImportTask struct {
	fileAccess    files.FileAccess
	id3Reader     files.ID3Reader
	mp3DbSearcher web.MP3DbSearch
	logger        logger.VerifyImportLogWriter
}

// NewVerifyImportTask - creates a new Instance of the VerifyImportTask
func NewVerifyImportTask(fileAccess files.FileAccess, id3Reader files.ID3Reader, mp3DbSearcher web.MP3DbSearch, logger logger.VerifyImportLogWriter) *VerifyImportTask {
	return &VerifyImportTask{fileAccess, id3Reader, mp3DbSearcher, logger}
}

// Execute - executes the task to import all MP3 files
func (task *VerifyImportTask) Execute(importDir string) (*model.VerifyImportStats, error) {
	filesToVerify, err := task.fileAccess.WalkDirectory(importDir)
	if err != nil {
		return nil, err
	}

	duplicatedFiles, failedFiles, err := task.startVerifyImport(filesToVerify)
	if err != nil {
		return nil, err
	}

	verifyImportStats := model.NewVerifyImportStats(importDir, filesToVerify, duplicatedFiles, failedFiles)

	task.logger.Log(verifyImportStats)

	return verifyImportStats, nil
}

func (task *VerifyImportTask) startVerifyImport(filesToVerify *[]string) (*[]string, *[]string, error) {
	countOfFilesToVerify := len(*filesToVerify)

	fmt.Printf("Start with verifying %v MP3 files...\n", countOfFilesToVerify)

	var duplicatedFiles []string
	var failedFiles []string

	bar := progress.NewInt(countOfFilesToVerify)

	for index, fileToVerify := range *filesToVerify {
		song, err := task.id3Reader.Read(fileToVerify)
		if err != nil {
			failedFiles = append(failedFiles, fileToVerify)
			continue
		}

		found, err := task.mp3DbSearcher.Search(song)
		if err != nil {
			return nil, nil, fmt.Errorf("Unexpected error while verifying songs for importing %v", err)
		}

		if found {
			duplicatedFiles = append(duplicatedFiles, fileToVerify)
		}

		updateVerifyImportProgress(bar, index)
	}

	updateVerifyImportProgress(bar, countOfFilesToVerify)
	fmt.Println()

	return &duplicatedFiles, &failedFiles, nil
}

func updateVerifyImportProgress(bar *progress.Bar, index int) {
	bar.ValueInt(index)
	bar.WriteTo(os.Stdout)
}
