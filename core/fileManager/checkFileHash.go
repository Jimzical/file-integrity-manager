package fileManager

import (
	"fmt"

	badger "github.com/dgraph-io/badger"
  bdgr "github.com/Jimzical/file-integrity-manager/pkg/badgerDB"
)

/*
CheckFileHash checks if the hash of a file already exists in the database.

This function checks if the hash of a file already exists in the database and adds it if it doesn't.

Parameters:
  - db: A pointer to the BadgerDB database.
  - filepath: The path of the file.
  - hash: The hash of the file.

Returns:
  - int: A code indicating the result of the hash check.
    NewEntryCode: constant indicating a new entry (value: 0)
    HashMatchCode: constant indicating a hash match (value: 1)
    HashMismatchCode: constant indicating a hash mismatch (value: 2)
    ErrorDuringHashCode: constant indicating an error during hash code (value: 3)
  - error: An error if the hash check fails.
*/
func CheckFileHash(db *badger.DB, filepath string, hash string) (int, error) {
    storedHash, err := bdgr.GetValueFromDB(db, filepath)
    if err != nil {
        return bdgr.ErrorDuringHashCode, err
    }

    if storedHash == "" {
        // Key doesn't exist, add it to the database
        err = bdgr.SetValueInDB(db, filepath, hash)
        if err != nil {
            return bdgr.ErrorDuringHashCode, fmt.Errorf("failed to add new entry: %v", err)
        }
        return bdgr.NewEntryCode, nil // Indicate that it's a new entry
    }

    if storedHash == hash {
        return bdgr.HashMatchCode, nil // Indicate that the hash matches
    }

    return bdgr.HashMismatchCode, nil // Indicate that the hash does not match
}
