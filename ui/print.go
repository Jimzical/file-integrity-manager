package ui

import (
	"fmt"
)

func Highlight(msg string) {
	// Apply the style to the print statement
	fmt.Print(HighlightStyle.Render(msg))
}

func Special(msg string) {
	// Apply the style to the print statement
	fmt.Print(SpecialStyle.Render(msg))
}

func Important(msg string) {
	// Apply the style to the print statement
	fmt.Print(ImportantStyle.Render(msg))
}

func Incorrect(msg string) {
	// Apply the style to the print statement
	fmt.Print(IncorrectStyle.Render(msg))
}

func Info(msg string) {
	// Apply the style to the print statement
	fmt.Print(InfoStyle.Render(msg))
}
