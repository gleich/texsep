package sort

import (
	"fmt"

	"github.com/Matt-Gleich/texsep/sort/files"
)

// CloneStructure ... Clone a structure
func CloneStructure() {
	files := files.ListRecursively()

	for _, file := range files {
		fmt.Printf("%#v\n", file)
	}
}
