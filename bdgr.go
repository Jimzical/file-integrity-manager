package main

import (
	"fmt"
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
)

func InitBadger() (*badger.DB, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    // Create /badger/filehash directory if it doesn't exist
    folderName := "/badger/filehash"
    dbDir := homeDir + folderName
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
        return Error, err
    }

    if storedHash == "" {
        // Key doesn't exist, add it to the database
        err = db.Update(func(txn *badger.Txn) error {
            return txn.Set([]byte(filepath), []byte(hash))
        })
        if err != nil {
            return Error, fmt.Errorf("failed to add new entry: %v", err)
        }

        return NewEntry, nil // Indicate that it's a new entry
    }

    if storedHash == hash {
        return HashMatch, nil // Indicate that the hash matches
    }

    return HashMismatch, nil // Indicate that the hash does not match
}
