package fileManager

import (
	"fmt"
	"sync"

	bdgr "github.com/Jimzical/file-integrity-manager/core/badgerDB"
	fileStructs "github.com/Jimzical/file-integrity-manager/core/models"
	badger "github.com/dgraph-io/badger"
	logs "github.com/Jimzical/file-integrity-manager/core/logs"
	status "github.com/Jimzical/file-integrity-manager/core/status"
)

// Deals with file hash and its management
//
// This function computes the hash of the file data and saves it to the database.
// It also checks if the hash of the file data already exists in the database.
//
// Parameters:
//   - filepathsChannel: A channel that receives the file paths to be hashed.
//   - db: A pointer to the BadgerDB database.
//   - wg: A pointer to the WaitGroup.
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

		statusType := status.GetStatus(result)

		switch statusType {
		case status.NEW_ENTRY:
			addedCount++
		case status.HASH_MATCH:
			matchCount++
		case status.HASH_MISMATCH:
			misMatchCount++
		}

		displayPath := logs.GetDisplayPath(filePath)
		rows = append(rows, []string{displayPath, statusType})
	}

	logs.PrintTable(rows)
}
