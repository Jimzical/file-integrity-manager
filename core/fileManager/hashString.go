package fileManager

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hashes a string using the SHA-256 algorithm.
//
// Parameters:
//   - data: The string to be hashed.
//
// Returns:
//   - string: The hexadecimal hash of the input string.
//
// Example usage:
//
//     hash := fileManager.hashString("Hello, World!")
//     fmt.Println("Hash:", hash)
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
