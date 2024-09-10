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

type fileStatus struct {
    filePath   string
    statusType string
}

/*
Deals with file hash and its management

This function computes the hash of the file data and saves it to the database.
It also checks if the hash of the file data already exists in the database.

Parameters:
  - filepathsChannel: A channel that receives the file paths to be hashed.
  - wg: A pointer to the WaitGroup.
*/
func (db *database) EncryptFiles(filepathsChannel <-chan fileStructs.FileInfo, wg *sync.WaitGroup) {
    defer wg.Done()

    var matchedRows [][]string
    var mismatchedRows [][]string
    var addedRows [][]string

    statusChannel := make(chan fileStatus)

    // Launch a goroutine to process files
    go db.processFiles(filepathsChannel, statusChannel)

    // Read from the statusChannel and update file counts
    for fileStatus := range statusChannel {
        switch fileStatus.statusType {
        case status.NEW_ENTRY:
            addedCount++
        case status.HASH_MATCH:
            matchCount++
        case status.HASH_MISMATCH:
            misMatchCount++
        }

        if configs.LOGGING_ENABLED {
            displayPath := logs.GetDisplayPath(fileStatus.filePath)
            switch fileStatus.statusType {
            case status.NEW_ENTRY:
                addedRows = append(addedRows, []string{displayPath, fileStatus.statusType})
            case status.HASH_MATCH:
                matchedRows = append(matchedRows, []string{displayPath, fileStatus.statusType})
            case status.HASH_MISMATCH:
                mismatchedRows = append(mismatchedRows, []string{displayPath, fileStatus.statusType})
            default:
                fmt.Printf("Unknown status type: %v\n", fileStatus.statusType)
            }
        }
    }

    if configs.LOGGING_ENABLED {
        ui.Success("Matched files\n")
        logs.PrintTable(matchedRows, status.HASH_MATCH)

        ui.Info("Added files\n")
        logs.PrintTable(addedRows, status.NEW_ENTRY)

        ui.Danger("Mismatched files\n")
        logs.PrintTable(mismatchedRows, status.HASH_MISMATCH)
    }
}

/*
Processes the files coming in from the filepathChannel and sends the status to the statusChannel to get classified and counted.

Parameters:
  - filepathsChannel: A channel that receives the file paths to be hashed.
  - statusChannel: A channel that sends the status of the file to be classified and counted.
*/
func (db *database) processFiles(filepathsChannel <-chan fileStructs.FileInfo, statusChannel chan<- fileStatus) {
    defer close(statusChannel)
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
        statusChannel <- fileStatus{filePath, statusType}
    }
}
