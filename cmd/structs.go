package main

import (
    "os"
    "time"
)

type FileInfo struct {
    FilePath string      `json:"filePath"`
    FileMode os.FileMode `json:"fileMode"`
    FileSize int64       `json:"fileSize"`
    ModTime  time.Time   `json:"modTime"`
}