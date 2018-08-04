package files

// FileAccess - defines the interface for the necessary file base functions
type FileAccess interface {
	ExistsDir(directoryPath string) bool
	ExistsFile(directoryPath string) bool
	WalkDirectory(rootDir string) (*[]string, error)
}
