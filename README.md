# File Integrity Manager

A command-line tool for managing file integrity, written entirely in Go.

## Description

File Integrity Manager ensures the integrity of files within a directory by generating and maintaining hashes for each file, including those in nested directories with a 3x performance boost compared to Python. 

It leverages the power of goroutines to process files concurrently, enhancing efficiency. The resulting hashes are stored in BadgerDB, a high-performance key-value store built in Go, designed for concurrent reads and writes leading in 3x performance boost compared to Python. Additionally, the tool employs Lipgloss to create a visually appealing user interface.

Users can check the integrity of their files by running the tool again, which compares the current hashes with those stored in the database and highlights any files that have been modified.

## Build

To build the tool, run the following command from the root directory of the project:

```bash
go build -o bin/file-integrity-manager.exe ./cmd/main.go
```

To run the executable, execute the following command:

```bash
./bin/file-integrity-manager.exe
```

## How to Run
To install the tool, follow these steps:
```bash
git clone https://github.com/Jimzical/file-integrity-manager.git
cd file-integrity-manager
go install
go mod tidy
```

To run the tool, execute the following commands:
```bash
cd /cmd
go run main.go
```

## Demo
![image](https://github.com/user-attachments/assets/923698b9-fcef-423a-b90b-ac7a1e09efd9)

