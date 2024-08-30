package fileManager

import (
	"fmt"
	"path/filepath"
	"strings"
)

func logHashResult(result int, filePath string) {
	// Split the file path into components
	components := strings.Split(filePath, string(filepath.Separator))
	displayPath := filePath

	if len(components) > 1 {
		// Get the last folder and the file name
		lastFolder := components[len(components)-2]
		fileName := components[len(components)-1]
		displayPath = "..." + string(filepath.Separator) + lastFolder + string(filepath.Separator) + fileName
	}

	switch result {
	case NewEntry:
		fmt.Printf("New entry added for %q\n", displayPath)
	case HashMatch:
		fmt.Printf("Hash match for %q\n", displayPath)
	case HashMismatch:
		fmt.Printf("Hash mismatch for %q\n", displayPath)
	default:
		fmt.Printf("Error checking hash for %q\n", displayPath)
	}
}
