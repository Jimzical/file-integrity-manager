package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

/*
Takes a string input from the user.
It prompts the user with a `highlighted` message and returns the input as a string.

Parameters:
  - msg: The message to display to the user.

Returns:
  - string: The input provided by the user.

Example usage:

    input := ui.InputFilePath("Enter your name:")
*/
func InputFilePath(msg string) string {
	defer fmt.Println()
	reader := bufio.NewReader(os.Stdin)

	Highlight(msg)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)             // Remove all leading and trailing whitespace, including \r and \n
	userInput = strings.Replace(userInput, "\"", "", -1) // remove quotes
	return userInput
}
