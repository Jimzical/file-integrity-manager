# File Integrity Manager

A command-line tool for managing file integrity, written entirely in Go.

## Description

File Integrity Manager ensures the integrity of files within a directory by generating and maintaining hashes for each file, including those in nested directories. It leverages the power of goroutines to process files concurrently, enhancing efficiency. The resulting hashes are stored in BadgerDB, a high-performance key-value store built in Go, designed for concurrent reads and writes. Additionally, the tool employs Lipgloss to create a visually appealing user interface.

Users can check the integrity of their files by running the tool again, which compares the current hashes with those stored in the database and highlights any files that have been modified.

## Build

To build the tool, run the following command from the root directory of the project:

```bash
go build -o bin/main ./cmd/main.go
```

## How to Run
To install the tool, follow these steps:
```bash
git clone https://github.com/Jimzical/file-integrity-manager.git
cd file-integrity-manager
go install
```

To run the tool, execute the following commands:
```bash
cd /cmd
go run main.go
```

## Usage
> Will Add this later



#TODO: Add a count for coorect map, apped and mishap map