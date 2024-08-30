package fileManager

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func getStatus(result int) string {
	switch result {
	case NewEntry:
		return "New Entry"
	case HashMatch:
		return "Hash Match"
	case HashMismatch:
		return "Hash Mismatch"
	default:
		return "Error Checking Hash"
	}
}

func getDisplayPath(filePath string) string {
	components := strings.Split(filePath, string(filepath.Separator))
	if len(components) > 1 {
		lastFolder := components[len(components)-2]
		fileName := components[len(components)-1]
		return "..." + string(filepath.Separator) + lastFolder + string(filepath.Separator) + fileName
	}
	return filePath
}

func printTable(rows [][]string) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("Display Path", "Status").
		Rows(rows...)

	fmt.Println(t)
}
