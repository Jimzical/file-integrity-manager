package fileManager

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	bdgr "github.com/Jimzical/file-integrity-manager/core/badgerDB"
	"github.com/Jimzical/file-integrity-manager/ui"

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
		fmt.Println("Folder could not be found, Please try again with the correct folder path")
		return
	}

	close(filepathsChannel)
	wg.Wait()

	ui.Special(fmt.Sprintf("\n\nMatched Files : %d\n", matchCount))
	fmt.Print("\033[1B\033[0G") // Move cursor down one line and to the beginning of the line
	ui.Incorrect(fmt.Sprintf("Mismatched Files : %d\n", misMatchCount))
	fmt.Print("\033[1B\033[0G") // Move cursor down one line and to the beginning of the line
	ui.Info(fmt.Sprintf("New Files : %d\n", addedCount))

}

func walkFolder(targetFolder string, filepathsChannel chan<- fileStructs.FileInfo) error {
	defer fmt.Print("\r")

	return filepath.Walk(targetFolder, func(file string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileCount++
		msg := fmt.Sprintf("\r\033[KFile %d: %s", fileCount, fileInfo.Name())

		ui.Special(msg)

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
