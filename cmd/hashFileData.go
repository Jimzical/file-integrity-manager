package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"

	bdgr "github.com/Jimzical/file-integrity-manager/internal/badgerDB"
	badger "github.com/dgraph-io/badger"
)

// Function to write paths to the file
func computeAndSaveFileHashes(filepathsChannel <-chan FileInfo, db *badger.DB , wg *sync.WaitGroup) {
	defer wg.Done()
 
	// Write the paths to the outputFile
	for file:= range filepathsChannel {
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

func hashString(data string) string {
    // Create a new SHA-256 hash object
    hash := sha256.New()
    
    // Write the string data to the hash object
    hash.Write([]byte(data))
    
    // Compute the hash and get the resulting byte slice
    hashBytes := hash.Sum(nil)
    
    // Convert the byte slice to a hexadecimal string
    hashString := hex.EncodeToString(hashBytes)
    
    return hashString
}

func logHashResult(result int, filePath string) {
	switch result {
	case NewEntry:
		fmt.Printf("New entry added for %q\n", filePath)
	case HashMatch:
		fmt.Printf("Hash match for %q\n", filePath)
	case HashMismatch:
		fmt.Printf("Hash mismatch for %q\n", filePath)
	default:
		fmt.Printf("Error checking hash for %q\n", filePath)
	}
}