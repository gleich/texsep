package sort

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsep/pkg/status"
)

// Move the files
func MoveFiles(filePaths []string) {
	if len(filePaths) != 0 {
		fmt.Println()
	}
	for _, filePath := range filePaths {
		var prefix string
		if strings.HasSuffix(filePath, ".tex") {
			latexPrefix := "./LaTeX/"
			if !strings.HasPrefix(filePath, latexPrefix) {
				prefix = latexPrefix
			}
		} else if strings.HasSuffix(filePath, ".pdf") {
			pdfPrefix := "./PDF/"
			if !strings.HasPrefix(filePath, pdfPrefix) {
				prefix = pdfPrefix
			}
		}
		fileFolderPathChunks := strings.Split(filePath, "/")
		fileFolderPath := prefix + strings.Join(fileFolderPathChunks[:len(fileFolderPathChunks)-1], "/") + "/"
		fileName := fileFolderPathChunks[len(fileFolderPathChunks)-1]

		_, err := os.Stat(fileFolderPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(fileFolderPath, 0700)
			if err != nil {
				statuser.Error("Failed to make folders", err, 1)
			}
		}

		err = os.Rename(filePath, fileFolderPath+fileName)
		if err != nil {
			statuser.Error("Failed to move file", err, 1)
		}
		status.Success("Moved " + filePath)
	}

}

// Getting all the tex and pdf files for the project
func Files() (filePaths []string) {
	// Getting files in project recursively
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

	return filePaths
}
