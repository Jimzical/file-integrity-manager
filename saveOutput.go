package main

import (
	"fmt"
	"os"
)

func saveFileWithTime(filePath string, fileHash string, file *os.File) {

	formatedData := fmt.Sprintf("%s,%s", fileHash, filePath)
    _, err := file.WriteString(string(formatedData) + "\n")

	if err != nil {
        fmt.Printf("Error writing to file %q: %v\n", file.Name(), err)
    }
}

