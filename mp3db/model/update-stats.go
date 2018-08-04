package model

// UpdateStats - Statistics about the last update
type UpdateStats struct {
	ImportStats
	UpdatedFilesCount int
}

// NewUpdateStats - Creates new UpdateStats with using the given arguments
func NewUpdateStats(
	rootDir string,
	filesToImport *[]string,
	importFilesCount int,
	updatedFilesCount int,
	failedFiles *[]string,
) *UpdateStats {

	totalCountOfFiles := len(*filesToImport)
	countOfFailedFiles := len(*failedFiles)

	return &UpdateStats{
		ImportStats: ImportStats{
			RootDirectory:      rootDir,
			ScannedFilesCount:  totalCountOfFiles,
			ImportedFilesCount: importFilesCount,
			FailedFilesCount:   countOfFailedFiles,
			FailedFiles:        *failedFiles,
		},
		UpdatedFilesCount: updatedFilesCount,
	}
}
