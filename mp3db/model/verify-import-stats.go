package model

// VerifyImportStats - Statistics about the last import verification
type VerifyImportStats struct {
	ImportDirectory      string
	VerifiedFilesCount   int
	DuplicatedFilesCount int
	DuplicatedFiles      []string
	FailedFilesCount     int
	FailedFiles          []string
}

// NewVerifyImportStats - Creates new VerifyImportStats
func NewVerifyImportStats(importDir string, filesToImport *[]string, duplicateFiles *[]string, failedFiles *[]string) *VerifyImportStats {
	totalCountOfFiles := len(*filesToImport)
	countOfDuplicatedFiles := len(*duplicateFiles)
	countOfFailedFiles := len(*failedFiles)

	return &VerifyImportStats{
		ImportDirectory:      importDir,
		VerifiedFilesCount:   totalCountOfFiles,
		DuplicatedFilesCount: countOfDuplicatedFiles,
		DuplicatedFiles:      *duplicateFiles,
		FailedFilesCount:     countOfFailedFiles,
		FailedFiles:          *failedFiles,
	}
}
