package fileManager

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	bdgr "github.com/Jimzical/file-integrity-manager/core/badgerDB"

	fileStructs "github.com/Jimzical/file-integrity-manager/core/models"
)

func TraverseFolder(targetFolder string) {
	var wg sync.WaitGroup
	filepathsChannel := make(chan fileStructs.FileInfo)

	db, err := bdgr.InitBadger()
	if err != nil {
		fmt.Printf("ErrorDuringHashCode initializing Badger: %v\n", err)
		return
	}
	defer db.Close()

	wg.Add(1)
	go ComputeAndSaveFileHashes(filepathsChannel, db, &wg)

	// Walk the folder and send file to the channel
	err = walkFolder(targetFolder, filepathsChannel)
	if err != nil {
		fmt.Printf("ErrorDuringHashCode walking the path %q: %v\n", targetFolder, err)
	}

	close(filepathsChannel)
	wg.Wait()
}

func walkFolder(targetFolder string, filepathsChannel chan<- fileStructs.FileInfo) error {
	return filepath.Walk(targetFolder, func(file string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// adding fileStructs.fileInfo to the channel to be hashed
		if !fileInfo.IsDir() {
			filepathsChannel <- fileStructs.FileInfo{
				FilePath: file,
				FileMode: fileInfo.Mode(),
				FileSize: fileInfo.Size(),
				ModTime:  fileInfo.ModTime(),
			}
		}

		return nil
	})
}
