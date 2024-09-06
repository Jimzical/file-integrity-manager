package main

import (
	"fmt"
	"time"

	fileManager "github.com/Jimzical/file-integrity-manager/core/fileManager"
	pkg "github.com/Jimzical/file-integrity-manager/pkg"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	ui.StartScreen()
	targetFolder := pkg.Input("Enter the target folder >")

	startTime := time.Now()
	fileManager.TraverseFolder(targetFolder)

	ui.Important(fmt.Sprintf("\n\nTime taken to complete: %v\n", time.Since(startTime)))

	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln()
}
