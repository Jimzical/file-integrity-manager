package ui

import (
	"fmt"
)

func Highlight(msg string) {
	// Apply the style to the print statement
	fmt.Print(highlight.Render(msg))
}

func Special(msg string) {
	// Apply the style to the print statement
	fmt.Print(special.Render(msg))
}

