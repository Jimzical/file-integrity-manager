package fileManager

import (
	"fmt"
	"sync"

	configs "github.com/Jimzical/file-integrity-manager/configs"
	logs "github.com/Jimzical/file-integrity-manager/core/logs"
	fileStructs "github.com/Jimzical/file-integrity-manager/core/models"
	status "github.com/Jimzical/file-integrity-manager/core/status"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

/*
Deals with file hash and its management

This function computes the hash of the file data and saves it to the database.
It also checks if the hash of the file data already exists in the database.

Parameters:
  - filepathsChannel: A channel that receives the file paths to be hashed.
  - db: A pointer to the BadgerDB database.
  - wg: A pointer to the WaitGroup.
*/
func (db *database) EncryptFile(filepathsChannel <-chan fileStructs.FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	var matchedRows [][]string
	var mismatchedRows [][]string
	var addedRows [][]string

	// Write the paths to the outputFile
	for file := range filepathsChannel {
		filePath := file.FilePath

		fileData := fmt.Sprintf("%s %v %d %v", filePath, file.FileMode, file.FileSize, file.ModTime)
		fileHash := hashString(fileData)

		result, err := db.CheckFileHash(filePath, fileHash)
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

		if configs.LOGGING_ENABLED {
			displayPath := logs.GetDisplayPath(filePath)
			switch statusType {
			case status.NEW_ENTRY:
				addedRows = append(addedRows, []string{displayPath, statusType})
			case status.HASH_MATCH:
				matchedRows = append(matchedRows, []string{displayPath, statusType})
			case status.HASH_MISMATCH:
				mismatchedRows = append(mismatchedRows, []string{displayPath, statusType})
			default:
				fmt.Printf("Unknown status type: %v\n", statusType)
			}
		}
	}
	if configs.LOGGING_ENABLED {
		ui.Special("Matched files\n")
		logs.PrintTable(matchedRows, status.HASH_MATCH)

		ui.Info("Added files\n")
		logs.PrintTable(addedRows, status.NEW_ENTRY)

		ui.Incorrect("Mismatched files\n")
		logs.PrintTable(mismatchedRows, status.HASH_MISMATCH)
	}
}
