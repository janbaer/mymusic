package files

// FileAccess - defines the interface for the necessary file base functions
type FileAccess interface {
	ExistsDir(directoryPath string) bool
	ExistsFile(filePath string) bool
	WalkDirectory(rootDir string) (*[]string, error)
	DeleteFile(filePath string) error
}
