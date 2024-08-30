package main

import (
	"fmt"

	fileManager "github.com/Jimzical/file-integrity-manager/internal/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	var targetFolder string
	fmt.Println("Starting")
	

	// targetFolder = "C:\\Personal\\Uni\\CS\\Golang\\file-integrity-manager\\tests"
	// TEST = C:\Personal\Uni\CS\Golang\file-integrity-manager\tests

	ui.StartScreen()

	targetFolder = fileManager.ReadTargetFolderPath()

	fileManager.TraverseFolder(targetFolder)

	fmt.Println("Done")
}
