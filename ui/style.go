package ui

import "github.com/charmbracelet/lipgloss"

// Style definitions.
var (
	// Generel
	Style = lipgloss.Style{}

	// General Stles.
	HighlightStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1bde4f")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1).
			Bold(true)

	SpecialStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#43BF6D")).
			Bold(true)

	ImportantStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Bold(true)

	IncorrectStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#de190b")).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4b13e8")).
			Bold(true)

	// Table.

	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("99")).
			Bold(true)

	BorderedTableStyle = lipgloss.NormalBorder()

	// Fonts.
	subtleFont  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	specialFont = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// Divider.
	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtleFont).
		String()

	// URL.
	url = lipgloss.NewStyle().Foreground(specialFont).Render

	// Title.
	descStyle = lipgloss.NewStyle().MarginTop(1)

	displayInfo = lipgloss.NewStyle().
			Align(lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtleFont)

	// Page.
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)
