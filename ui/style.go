package ui

import "github.com/charmbracelet/lipgloss"

// Style definitions.
var (

	// General.
	highlight = lipgloss.NewStyle().
			Background(lipgloss.Color("#1bde4f")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1).
			Bold(true)
	subtle  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	special = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Title.
	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Page.
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)
