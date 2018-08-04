package logger

import (
	"fmt"
	"os"
	"path"

	"github.com/janbaer/mp3db/model"
)

// UpdateLogWriter - defines function to write the updatestats to a log file
type UpdateLogWriter interface {
	WriteLog(importState *model.UpdateStats) error
}

// UpdateLogLogger - implements the interface UpdateLogWriter
type UpdateLogLogger struct {
	LogDirPath string
}

// NewUpdateLogLogger - Creates a new UpdateLogLogger
func NewUpdateLogLogger(logDirPath string) *UpdateLogLogger {
	return &UpdateLogLogger{LogDirPath: logDirPath}
}

// WriteLog - Creates a logfile with stats about the last update
func (logger UpdateLogLogger) WriteLog(updateStats *model.UpdateStats) error {
	updateLogFilepath := path.Join(logger.LogDirPath, "update.log")
	os.Remove(updateLogFilepath)

	file, err := os.Create(updateLogFilepath)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(
		fmt.Sprintf("Update start from %s with %d files total, %d files imported, %d files updated, %d files failed\n",
			updateStats.RootDirectory,
			updateStats.ScannedFilesCount,
			updateStats.ImportedFilesCount,
			updateStats.UpdatedFilesCount,
			updateStats.FailedFilesCount,
		))

	file.WriteString("-------------------------------------------------------\n")
	for _, failedFile := range updateStats.FailedFiles {
		file.WriteString(fmt.Sprintf("%s\n", failedFile))
	}

	return nil
}
