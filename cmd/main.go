package main

import (
	"fmt"
	"time"

	fileManager "github.com/Jimzical/file-integrity-manager/core/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
	basics "github.com/Jimzical/file-integrity-manager/pkg/basics"
)

func main() {
	ui.StartScreen()
	targetFolder := basics.Input("Enter the target folder >")

	startTime := time.Now()
	fileManager.TraverseFolder(targetFolder)

	ui.Important(fmt.Sprintf("\n\nTime taken to complete: %v\n", time.Since(startTime)))

}
