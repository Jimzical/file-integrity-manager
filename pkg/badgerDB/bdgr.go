package badgerDB

import (
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
	configs "github.com/Jimzical/file-integrity-manager/configs"
)

/*
Initialize BadgerDB

This function initializes BadgerDB and returns a pointer to the database.

Returns:
  - *badger.DB: A pointer to the BadgerDB database.
  - error: An error if the database initialization fails.

Example usage:

    db, err := InitBadger()
    if err != nil {
        log.Fatal(err)
    }
*/
func InitBadger() (*badger.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create /badger/filehash directory if it doesn't exist
	folderName := configs.BadgerFolderName
	dbDir := homeDir + "/badger/" + folderName // Fixed path concatenation
	err = os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	opt := badger.DefaultOptions(dbDir)
	opt.Logger = nil // Disable Badger's default logger
	opt.Truncate = true
  opt.SyncWrites = false

	db, err := badger.Open(opt)

	return db, err
}

/*
GetValueFromDB retrieves a value from the Badger database.

Parameters:
  - db: A pointer to the BadgerDB database.
  - key: The key to retrieve.

Returns:
  - string: The value of the key.
  - error: An error if the key-value pair cannot be retrieved.
*/
func GetValueFromDB(db *badger.DB, key string) (string, error) {
    var value string
    err := db.View(func(txn *badger.Txn) error {
        item, err := txn.Get([]byte(key))
        if err != nil {
            if err == badger.ErrKeyNotFound {
                return nil // Key doesn't exist, no error
            }
            return err
        }

        return item.Value(func(val []byte) error {
            value = string(val)
            return nil
        })
    })
    return value, err
}

/*
SetValueInDB sets a key-value pair in the Badger database.

Parameters:
  - db: A pointer to the BadgerDB database.
  - key: The key to set.
  - value: The value to set.

Returns:
  - error: An error if the key-value pair cannot be set.
*/
func SetValueInDB(db *badger.DB, key, value string) error {
    return db.Update(func(txn *badger.Txn) error {
        return txn.Set([]byte(key), []byte(value))
    })
}


