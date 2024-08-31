package main

import (
	"fmt"
	"time"

	fileManager "github.com/Jimzical/file-integrity-manager/core/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	ui.StartScreen()

	startTime := time.Now()
	// targetFolder := fileManager.Input("Enter the target folder >")
	// BorderForeground(lipgloss.Color("63"))
	// TEMP
	targetFolder := "C:\\Personal\\Uni\\CS\\Golang\\file-integrity-manager\\tests"
	fileManager.TraverseFolder(targetFolder)

	ui.Important(fmt.Sprintf("\n\nTime taken to complete: %v\n", time.Since(startTime)))

}
