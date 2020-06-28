package sort

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsep/status"
)

// MoveFiles ... Move the files
func MoveFiles(filePaths []string) {
	fmt.Println("")
	for _, filePath := range filePaths {
		var prefix string
		if strings.HasSuffix(filePath, ".tex") {
			prefix = "./tex/"
		} else if strings.HasSuffix(filePath, ".pdf") {
			prefix = "./pdf/"
		}
		fileFolderPathChunks := strings.Split(filePath, "/")
		fileFolderPath := prefix + strings.Join(fileFolderPathChunks[:len(fileFolderPathChunks)-1], "/") + "/"
		fileName := fileFolderPathChunks[len(fileFolderPathChunks)-1]

		_, err := os.Stat(fileFolderPath)
		if os.IsNotExist(err) {
			os.MkdirAll(fileFolderPath, 0700)
		}

		err = os.Rename(filePath, fileFolderPath+fileName)
		if err != nil {
			statuser.Error("Failed to move file", err, 1)
		}
		status.Success("Moved " + filePath)
	}

}

// Files ... Getting all the tex and pdf files for the project
func Files() (filePaths []string) {
	// Getting files in project recursively:
	allFiles := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		allFiles = append(allFiles, path)
		return nil
	})
	if err != nil {
		statuser.Error("Failed to get all tex and pdf files", err, 1)
	}

	for _, filePath := range allFiles {
		if strings.HasSuffix(filePath, ".tex") || strings.HasSuffix(filePath, ".pdf") {
			filePaths = append(filePaths, filePath)
		}
	}

	status.Success("Successfully got all tex and pdf files")

	return filePaths
}
