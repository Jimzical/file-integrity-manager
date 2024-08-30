package main

import (
	"fmt"

	fileManager "github.com/Jimzical/file-integrity-manager/internal/fileManager"
	ui "github.com/Jimzical/file-integrity-manager/ui"
)

func main() {
	fmt.Println("Starting")

	targetFolder := "C:\\Personal\\Uni\\CS\\Golang\\file-integrity-manager\\tests"

	ui.StartScreen()

	fileManager.TraverseFolder(targetFolder)

	fmt.Println("Done")
}
