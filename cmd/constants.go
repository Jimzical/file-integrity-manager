package main

// For Files
const (
	targetFolder = "C:\\Personal\\Uni\\CS\\Golang\\file-integrity-manager\\testingFolder"
)

// For Badger
const (
	Error        = iota - 1 // -1
	NewEntry                // 0
	HashMatch               // 1
	HashMismatch            // 2
)
