package logger

import (
	"fmt"
	"os"

	"github.com/janbaer/mp3db/model"
)

// ImportLogWriter - defines function to write to importstats to a log file
type ImportLogWriter interface {
	WriteLog(importState *model.ImportStats) error
}

// ImportLogLogger - implements the interface ImportLogWriter
type ImportLogLogger struct {
}

// WriteLog - Creates a logfile with stats about the last import
func (logger ImportLogLogger) WriteLog(importStats *model.ImportStats) error {
	importLogFilepath := "./import.log"
	os.Remove(importLogFilepath)

	file, err := os.Create(importLogFilepath)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(
		fmt.Sprintf("Import start from %s with %d files total, %d files imported, %d files failed\n",
			importStats.RootDirectory,
			importStats.ScannedFilesCount,
			importStats.ImportedFilesCount,
			importStats.FailedFilesCount,
		))

	file.WriteString("-------------------------------------------------------\n")
	for _, failedFile := range importStats.FailedFiles {
		file.WriteString(fmt.Sprintf("%s\n", failedFile))
	}

	return nil
}
