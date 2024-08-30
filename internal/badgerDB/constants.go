package badgerDB

const (
	Error        = iota - 1 // -1
	NewEntry                // 0
	HashMatch               // 1
	HashMismatch            // 2
)
