package ui

import (
	"fmt"
)

/*
Highlight the message with a Green style.

Parameters:
  - msg: The message to print.
*/
func Highlight(msg interface{}) {
    // Apply the style to the print statement
    fmt.Print(HighlightStyle.Render(msg.(string)))
}

/*
Style the message with a Green style.

Parameters:
  - msg: The message to print.
*/
func Special(msg interface{}) {
    // Apply the style to the print statement
    fmt.Print(SpecialStyle.Render(msg.(string)))
}

/*
Style the message with a White style.

Parameters:
  - msg: The message to print.
*/
func Important(msg interface{}) {
    // Apply the style to the print statement
    fmt.Print(ImportantStyle.Render(msg.(string)))
}

/*
Style the message with a Red style.

Parameters:
  - msg: The message to print.
*/
func Incorrect(msg interface{}) {
    // Apply the style to the print statement
    fmt.Print(IncorrectStyle.Render(msg.(string)))
}

/*
Style the message with a Yellow style.

Parameters:
  - msg: The message to print.
*/
func Info(msg interface{}) {
    // Apply the style to the print statement
    fmt.Print(InfoStyle.Render(msg.(string)))
}