package fileManager

import badger "github.com/dgraph-io/badger"

type database struct {
	db *badger.DB
}
