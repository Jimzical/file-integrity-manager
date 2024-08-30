package fileManager

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func ReadTargetFolderPath() string {
	reader := bufio.NewReader(os.Stdin)

	// Define the style
	style := lipgloss.NewStyle().
		Background(lipgloss.Color("#1bde4f")).
		Foreground(lipgloss.Color("#000000")).
		Padding(0, 1).
		Bold(true)

	// Apply the style to the print statement
	fmt.Print(style.Render("Enter the target folder path: "))

	targetFolder, _ := reader.ReadString('\n')
	targetFolder = strings.TrimSpace(targetFolder)             // Remove all leading and trailing whitespace, including \r and \n
	targetFolder = strings.Replace(targetFolder, "\"", "", -1) // remove quotes
	return targetFolder
}
