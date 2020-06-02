package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/texsep/status"
)

// ListRecursively ... Get a list of all files in the current and all sub-directories
func ListRecursively() []string {
	files := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		status.Error(err, "Failed to get all files")
	}

	cleanedFiles := []string{}
	for _, dirtyFile := range files {
		if strings.HasSuffix(dirtyFile, ".tex") || strings.HasSuffix(dirtyFile, ".pdf") {
			cleanedFiles = append(cleanedFiles, dirtyFile)
		}
	}

	return cleanedFiles
}
