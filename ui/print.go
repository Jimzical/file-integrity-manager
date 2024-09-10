package ui

import (
	"fmt"
)

// This code can be improved later by structuing the code differently


/*
Highlight the message with a Green style.

Parameters:
  - msg: The message to print.
*/
func Highlight(msg any) {
    // Apply the style to the print statement
    fmt.Print(HighlightStyle.Render(msg.(string)))
}

/*
Style the message with a Green style.

Parameters:
  - msg: The message to print.
*/
func Special(msg any) {
    // Apply the style to the print statement
    fmt.Print(SpecialStyle.Render(msg.(string)))
}

/*
Style the message with a White style.

Parameters:
  - msg: The message to print.
*/
func Important(msg any) {
    // Apply the style to the print statement
    fmt.Print(ImportantStyle.Render(msg.(string)))
}

/*
Style the message with a Red style.

Parameters:
  - msg: The message to print.
*/
func Incorrect(msg any) {
    // Apply the style to the print statement
    fmt.Print(IncorrectStyle.Render(msg.(string)))
}

/*
Style the message with a Yellow style.

Parameters:
  - msg: The message to print.
*/
func Info(msg any) {
    // Apply the style to the print statement
    fmt.Print(InfoStyle.Render(msg.(string)))
}

/*
Highlight the message with a Green style using formatted input.

Parameters:
  - format: The format string.
  - args: The arguments to format.
*/
func HighlightF(format string, args ...any) {
    // Apply the style to the formatted print statement
    fmt.Print(HighlightStyle.Render(fmt.Sprintf(format, args...)))
}

/*
Style the message with a Green style using formatted input.

Parameters:
  - format: The format string.
  - args: The arguments to format.
*/
func SpecialF(format string, args ...any) {
    // Apply the style to the formatted print statement
    fmt.Print(SpecialStyle.Render(fmt.Sprintf(format, args...)))
}

/*
Style the message with a White style using formatted input.

Parameters:
  - format: The format string.
  - args: The arguments to format.
*/
func ImportantF(format string, args ...any) {
    // Apply the style to the formatted print statement
    fmt.Print(ImportantStyle.Render(fmt.Sprintf(format, args...)))
}

/*
Style the message with a Red style using formatted input.

Parameters:
  - format: The format string.
  - args: The arguments to format.
*/
func IncorrectF(format string, args ...any) {
    // Apply the style to the formatted print statement
    fmt.Print(IncorrectStyle.Render(fmt.Sprintf(format, args...)))
}

/*
Style the message with a Yellow style using formatted input.

Parameters:
  - format: The format string.
  - args: The arguments to format.
*/
func InfoF(format string, args ...any) {
    // Apply the style to the formatted print statement
    fmt.Print(InfoStyle.Render(fmt.Sprintf(format, args...)))
}