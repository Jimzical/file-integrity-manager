package fileManager

import (
	"fmt"
	"os"
	"sync"

	"github.com/iafan/cwalk"

	fileStructs "github.com/Jimzical/file-integrity-manager/core/models"
	bdgr "github.com/Jimzical/file-integrity-manager/pkg/badgerDB"
	"github.com/Jimzical/file-integrity-manager/pkg/basics"
	ui "github.com/Jimzical/file-integrity-manager/ui"
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

	database := &database{db: db}

	wg.Add(1)
	go database.FileManager(filepathsChannel, &wg)

	// Walk the folder and send file to the channel
	err = walkFolder(targetFolder, filepathsChannel)
	if err != nil {
		fmt.Println("Folder could not be found, Please try again with the correct folder path")
		return
	}
	close(filepathsChannel)

	basics.ClearAndPrint("")

	wg.Wait()

	ui.Success(fmt.Sprintf("\n\nMatched Files : %d\n", matchCount))
	fmt.Println()
	ui.Danger(fmt.Sprintf("Mismatched Files : %d\n", misMatchCount))
	fmt.Println()
	ui.Info(fmt.Sprintf("New Files : %d\n", addedCount))

}

/*
walkFolder walks the target folder and sends the file paths to the filepathsChannel.

Parameters:
  - targetFolder: The folder to walk.
  - filepathsChannel: A channel that receives the file paths to be hashed.

Returns:
  - error: An error if the folder walk fails.
*/
func walkFolder(targetFolder string, filepathsChannel chan<- fileStructs.FileInfo) error {
	defer fmt.Print("\r")
	var fileCount int

	return cwalk.Walk(targetFolder, func(file string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileCount++
		// Clear the line and print the file name every 100 files
		if fileCount%100 == 0 {
			msg := basics.ClearAndSprintf("File %d: %s", fileCount, fileInfo.Name())
			ui.Success(msg)
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
