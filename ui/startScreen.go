package ui

// This example demonstrates various Lip Gloss style and layout features.

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func StartScreen() {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))

	doc := strings.Builder{}

	// Title
	{
		var (
			title strings.Builder
		)

		desc := lipgloss.JoinVertical(lipgloss.Left,
			descStyle.Render(lipgloss.NewStyle().PaddingRight(10).Foreground(specialFont).Background(lipgloss.Color("#2e302f")).Render(`
		
		___ _ _       ___     _                  __          __  __                             
		| __(_| |___  |_ _|_ _| |_ ___ __ _ _ _(_| |_ _  _  |  \/  |__ _ _ _  __ _ __ _ ___ _ _ 
		| _|| | / -_)  | || ' |  _/ -_/ _' | '_| |  _| || | | |\/| / _' | ' \/ _' / _' / -_| '_|
		|_| |_|_\___| |___|_||_\__\___\__, |_| |_|\__|\_, | |_|  |_\__,_|_||_\__,_\__, \___|_|  
									|___/           |__/                        |___/         

			`)),
			displayInfo.Render("From"+divider+url("https://github.com/Jimizical/file-integrity-manager")),
		)

		row := lipgloss.JoinHorizontal(lipgloss.Top, title.String(), desc)
		doc.WriteString(row + "\n\n")
	}

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	fmt.Println(docStyle.Render(doc.String()))
}
