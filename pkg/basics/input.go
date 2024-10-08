package basics

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Takes a string input from the user.
It prompts the user with a message and returns the input as a string.

Parameters:
  - msg: The message to display to the user.

Returns:
  - string: The input provided by the user.

Example usage:

    input := basics.InputFilePath("Enter your name:")
*/
func InputFilePath(msg string) string {
	defer fmt.Println()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(msg)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)             // Remove all leading and trailing whitespace, including \r and \n
	userInput = strings.Replace(userInput, "\"", "", -1) // remove quotes
	return userInput
}
