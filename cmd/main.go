package main

import (
	"fmt"

	fileManager "github.com/Jimzical/file-integrity-manager/internal/fileManager"
)

func main() {
	fmt.Println("Starting")

	targetFolder := "C:\\Personal\\Uni\\CS\\Golang\\file-integrity-manager\\tests\\testingFolder"

	fileManager.TraverseFolder(targetFolder)

	fmt.Println("Done")
}
