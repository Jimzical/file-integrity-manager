package main

import (
	"fmt"
	"time"

	fileManager "github.com/Jimzical/file-integrity-manager/core/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	ui.StartScreen()
	targetFolder := fileManager.Input("Enter the target folder >")

	startTime := time.Now()
	fileManager.TraverseFolder(targetFolder)

	ui.Important(fmt.Sprintf("\n\nTime taken to complete: %v\n", time.Since(startTime)))

	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln()
}
