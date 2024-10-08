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

// getStatus returns the status of the file hash check
//
// Parameters:
//   - result: The result code of the hash check
//
// Returns:
//   - string: The status of the hash check {NEW_ENTRY |HASH_MATCH | HASH_MISMATCH| ERROR_CHECKING_HASH}
func GetStatus(result int) string {
	switch result {
	case NewEntryCode:
		return NEW_ENTRY
	case HashMatchCode:
		return HASH_MATCH
	case HashMismatchCode:
		return HASH_MISMATCH
	default:
		return ERROR_CHECKING_HASH
	}
}
