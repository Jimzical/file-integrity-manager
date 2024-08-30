package fileManager

import (
	"fmt"
	"sync"

	badger "github.com/dgraph-io/badger"

	bdgr "github.com/Jimzical/file-integrity-manager/internal/badgerDB"
	fileStructs "github.com/Jimzical/file-integrity-manager/internal/models"
)

// Function to write paths to the file
func ComputeAndSaveFileHashes(filepathsChannel <-chan fileStructs.FileInfo, db *badger.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	// Write the paths to the outputFile
	for file := range filepathsChannel {
		filePath := file.FilePath

		fileData := fmt.Sprintf("%s %v %d %v", filePath, file.FileMode, file.FileSize, file.ModTime)
		fileHash := hashString(fileData)

		result, err := bdgr.CheckFileHash(db, filePath, fileHash)
		if err != nil {
			fmt.Printf("Error checking file hash %q: %v\n", filePath, err)
			continue
		}

		logHashResult(result, filePath)

	}
}