package main

// For Files
const (
	targetFolder = "testingFolder"
	storageFile  = "output/output.csv"
)

// For Badger
const (
    Error = iota - 1 // -1
    NewEntry         // 0
    HashMatch        // 1
    HashMismatch     // 2
)