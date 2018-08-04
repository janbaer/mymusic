package files

import (
	"os/user"
	"strings"
)

// ExpandHomeDirIfNeeded - This function checks if the given path is starting with the ~ homedir
// alias and returns the expanded path instead. If not it just returns the given path
func ExpandHomeDirIfNeeded(path string) string {
	if strings.Index(path, "~") == 0 {
		user, _ := user.Current()
		homeDir := user.HomeDir
		path = strings.Replace(path, "~", homeDir, 1)
	}

	return path
}
