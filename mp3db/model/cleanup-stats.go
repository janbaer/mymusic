package model

// CleanupStats - Statistics about the last cleanup
type CleanupStats struct {
	TotalCountOfSongs        int
	TotalCountOfDeletedSongs int
	DeletedSongs             []Song
}

// NewCleanupStats - Creates new ImportStats with using the given arguments
func NewCleanupStats(totalCountOfSongs int, totalCountOfDeletedSongs int, deletedSongs *[]Song) *CleanupStats {
	return &CleanupStats{
		TotalCountOfSongs:        totalCountOfSongs,
		TotalCountOfDeletedSongs: totalCountOfDeletedSongs,
		DeletedSongs:             *deletedSongs,
	}
}
