package basics

import (
    "fmt"
)

/*
Clear the current line and print the message.

Parameters:
  - msg: The message to print.

Example:
	basics.ClearAndPrint("Processing %d file...", fileCount)
*/
func ClearAndPrint(msg any) {
    fmt.Print("\r\033[K")
    fmt.Print(msg)
}

/*
Clear the current line and return the formatted message.

Parameters:
  - format: The format string.
  - a: The arguments to format.

Returns:
  - string: The formatted message.

Example:
  msg := basics.ClearAndSprintf("File %d: %s", fileCount, fileInfo.Name())
*/
func ClearAndSprintf(format string, a ...any) string {
    msg := fmt.Sprintf(format, a...)
    return "\r\033[K" + msg
}