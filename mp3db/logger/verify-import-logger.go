package logger

import (
	"fmt"

	"github.com/janbaer/mp3db/model"
)

// VerifyImportLogWriter - defines function to log the VerifyImportStats
type VerifyImportLogWriter interface {
	Log(verifyImportStats *model.VerifyImportStats) error
}

// VerifyImportLogLogger - Implements the Logger for the VerifyImportStats
type VerifyImportLogLogger struct {
}

// NewVerifyImportLogLogger - Creates a new VerifyImportLogLogger
func NewVerifyImportLogLogger() *VerifyImportLogLogger {
	return &VerifyImportLogLogger{}
}

// Log - Logs the VerifyImportStats
func (logger VerifyImportLogLogger) Log(verifyImportStats *model.VerifyImportStats) error {

	if verifyImportStats.DuplicatedFilesCount == 0 {
		return nil
	}

	fmt.Println("The following duplicates were found:")
	for _, fileName := range verifyImportStats.DuplicatedFiles {
		fmt.Printf("%s\n", fileName)
	}

	return nil
}
