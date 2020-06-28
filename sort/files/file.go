package files

import (
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/statuser/v2"
)

// ListRecursively ... Get a list of all files in the current and all sub-directories
func ListRecursively() []string {
	files := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		statuser.Error("Failed to get all files", err, 1)
	}
	return files
}
