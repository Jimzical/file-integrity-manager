package badgerDB

import (
	"fmt"
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
)
// Initialize BadgerDB
//
// This function initializes BadgerDB and returns a pointer to the database.
//
// Returns:
//   - *badger.DB: A pointer to the BadgerDB database.
//   - error: An error if the database initialization fails.
//
// Example usage:
//
//     db, err := InitBadger()
//     if err != nil {
//         log.Fatal(err)
//     }
func InitBadger() (*badger.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create /badger/filehash directory if it doesn't exist
	folderName := dbFolderName
	dbDir := homeDir + "/badger/" + folderName // Fixed path concatenation
	err = os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	opt := badger.DefaultOptions(dbDir)
	opt.Logger = nil // Disable Badger's default logger
	opt.Truncate = true
	db, err := badger.Open(opt)

	return db, err
}

// CheckFileHash checks if the hash of a file already exists in the database.
//
// This function checks if the hash of a file already exists in the database and adds it if it doesn't.
//
// Parameters:
//   - db: A pointer to the BadgerDB database.
//   - filepath: The path of the file.
//   - hash: The hash of the file.
//
// Returns:
//   int: A code indicating the result of the hash check.
//     NewEntryCode: constant indicating a new entry (value: 0)
//     HashMatchCode: constant indicating a hash match (value: 1)
//     HashMismatchCode: constant indicating a hash mismatch (value: 2)
//     ErrorDuringHashCode: constant indicating an error during hash code (value: 3)
//   error: An error if the hash check fails.
func CheckFileHash(db *badger.DB, filepath string, hash string) (int, error) {
	var storedHash string
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(filepath))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil // Key doesn't exist, no error
			}
			return err
		}

		return item.Value(func(val []byte) error {
			storedHash = string(val)
			return nil
		})
	})

	if err != nil {
		return ErrorDuringHashCode, err
	}

	if storedHash == "" {
		// Key doesn't exist, add it to the database
		err = db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(filepath), []byte(hash))
		})
		if err != nil {
			return ErrorDuringHashCode, fmt.Errorf("failed to add new entry: %v", err)
		}

		return NewEntryCode, nil // Indicate that it's a new entry
	}

	if storedHash == hash {
		return HashMatchCode, nil // Indicate that the hash matches
	}

	return HashMismatchCode, nil // Indicate that the hash does not match
}
