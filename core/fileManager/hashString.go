package fileManager

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashString(data string) string {
	// Create a new SHA-256 hash object
	hash := sha256.New()

	// Write the string data to the hash object
	hash.Write([]byte(data))

	// Compute the hash and get the resulting byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
