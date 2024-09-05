package status

// Match Codes
const (
	ErrorDuringHashCode = iota - 1 // -1
	NewEntryCode                   // 0
	HashMatchCode                  // 1
	HashMismatchCode               // 2
)

// Match Strings
const (
	NEW_ENTRY           = "New Entry"
	HASH_MATCH          = "Hash Match"
	HASH_MISMATCH       = "Hash Mismatch"
	ERROR_CHECKING_HASH = "ErrorDuringHashCode Checking Hash"
)
