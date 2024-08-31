package fileManager

import (
	"fmt"
	"sync"

	bdgr "github.com/Jimzical/file-integrity-manager/core/badgerDB"
	fileStructs "github.com/Jimzical/file-integrity-manager/core/models"
	badger "github.com/dgraph-io/badger"
)

// Function to write paths to the file
func ComputeAndSaveFileHashes(filepathsChannel <-chan fileStructs.FileInfo, db *badger.DB, wg *sync.WaitGroup) {
	defer wg.Done()

	var rows [][]string

	// Write the paths to the outputFile
	for file := range filepathsChannel {
		filePath := file.FilePath

		fileData := fmt.Sprintf("%s %v %d %v", filePath, file.FileMode, file.FileSize, file.ModTime)
		fileHash := hashString(fileData)

		result, err := bdgr.CheckFileHash(db, filePath, fileHash)
		if err != nil {
			fmt.Printf("ErrorDuringHashCode checking file hash %q: %v\n", filePath, err)
			continue
		}

		status := getStatus(result)

		switch status {
		case NEW_ENTRY:
			addedCount++
		case HASH_MATCH:
			matchCount++
		case HASH_MISMATCH:
			misMatchCount++
		}

		displayPath := getDisplayPath(filePath)
		rows = append(rows, []string{displayPath, status})
	}

	printTable(rows)
}
