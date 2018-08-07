package files

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
)

// FileSystem - Provides the real implementations for FileAccess
type FileSystem struct {
}

// ExistsDir - werifies if the given directoryPath exists and if it's a directory
func (fileSystem FileSystem) ExistsDir(directoryPath string) bool {
	fileInfo, err := os.Stat(directoryPath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// ExistsFile - Verifies if the given file exists
func (fileSystem FileSystem) ExistsFile(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

// WalkDirectory - walks through all directories beginning from the rootDirectory
// and returns all found MP3 files
func (fileSystem FileSystem) WalkDirectory(rootDir string) (*[]string, error) {
	if !fileSystem.ExistsDir(rootDir) {
		return nil, fmt.Errorf("the directory %s you passed, not exists", rootDir)
	}

	s := configureAndStartSpinner(rootDir)
	defer s.Stop()

	var filesToImport []string

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename := info.Name()
		if !info.IsDir() && filepath.Ext(filename) == ".mp3" {
			filesToImport = append(filesToImport, path)
		}

		return nil
	})

	fmt.Println()
	return &filesToImport, nil
}

// DeleteFile - deletes the file with the passed filePath
func (fileSystem FileSystem) DeleteFile(filePath string) error {
	if fileSystem.ExistsFile(filePath) {
		return os.Remove(filePath)
	}
	return nil
}

func configureAndStartSpinner(rootDir string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Start with importing %s, please wait... ", rootDir)
	s.Color("red")
	s.Start()
	return s
}
