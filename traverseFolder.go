package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func traverseFolder(targetFolder string) {
	var wg sync.WaitGroup
	filepathsChannel := make(chan FileInfo)

	db, err := InitBadger()
	if err != nil {
		fmt.Printf("Error initializing Badger: %v\n", err)
		return
	}
	defer db.Close()

	wg.Add(1)
	go computeAndSaveFileHashes(filepathsChannel, db, &wg)

    // // or
    // go computeAndVerifyHashes(filepathsChannel, &wg)

    /*
        So here the idea would be to just have a swtich case by haveing a flag for mode
        and then based on the flag we can call the respective function

        Hashing mechanism can be same for both the cases, only the verification part will be different
        as rahter thatn the saveFileWithTime()
        we will have a function to read the file and then compare the hash
        and print out any mismatches, else print out the success message 
    */



	// Walk the folder and send file to the channel
	err = walkFolder(targetFolder, filepathsChannel)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", targetFolder, err)
	}

	close(filepathsChannel)
	wg.Wait()
}

func walkFolder(targetFolder string, filepathsChannel chan<- FileInfo) error {
	return filepath.Walk(targetFolder, func(file string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// adding fileInfo to the channel to be hashed
		if !fileInfo.IsDir() {
			filepathsChannel <- FileInfo{
				FilePath: file,
				FileMode: fileInfo.Mode(),
				FileSize: fileInfo.Size(),
				ModTime:  fileInfo.ModTime(),
			}
		}

		return nil
	})
}
