package badgerDB

const (
	ErrorDuringHashCode = iota - 1 // -1
	NewEntryCode                   // 0
	HashMatchCode                  // 1
	HashMismatchCode               // 2
)
