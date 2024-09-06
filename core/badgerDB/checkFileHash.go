package badgerDB

import (
	"fmt"

	badger "github.com/dgraph-io/badger"
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
    storedHash, err := GetValueFromDB(db, filepath)
    if err != nil {
        return ErrorDuringHashCode, err
    }

    if storedHash == "" {
        // Key doesn't exist, add it to the database
        err = SetValueInDB(db, filepath, hash)
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
