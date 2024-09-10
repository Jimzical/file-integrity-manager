package main

import (
	"time"

	fileManager "github.com/Jimzical/file-integrity-manager/core/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	ui.StartScreen()
	targetFolder := ui.InputFilePath("Enter the target folder >")

	startTime := time.Now()
	fileManager.TraverseFolder(targetFolder)


	ui.ImportantF("\n\nTime taken to complete: %v\n", time.Since(startTime))
}
