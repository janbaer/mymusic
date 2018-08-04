package model

// ImportStats - Statistics about the last import
type ImportStats struct {
	RootDirectory      string
	ScannedFilesCount  int
	ImportedFilesCount int
	FailedFilesCount   int
	FailedFiles        []string
}

// NewImportStats - Creates new ImportStats with using the given arguments
func NewImportStats(rootDir string, filesToImport *[]string, failedFiles *[]string) *ImportStats {
	totalCountOfFiles := len(*filesToImport)
	countOfFailedFiles := len(*failedFiles)

	return &ImportStats{
		RootDirectory:      rootDir,
		ScannedFilesCount:  totalCountOfFiles,
		ImportedFilesCount: totalCountOfFiles - countOfFailedFiles,
		FailedFilesCount:   countOfFailedFiles,
		FailedFiles:        *failedFiles,
	}
}
