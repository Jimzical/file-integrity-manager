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
	case NewEntryCode:
		return NEW_ENTRY
	case HashMatchCode:
		return HASH_MATCH
	case HashMismatchCode:
		return HASH_MISMATCH
	default:
		return ERROR_CHECKING_HASH
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
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				// Style for the header row
				return lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
			}
			if rows[row-1][STATUS_COL] == HASH_MISMATCH {
				// return lipgloss.NewStyle().Background(lipgloss.Color("#de190b")).Foreground(lipgloss.Color("#ffffff"))
				return lipgloss.NewStyle().Foreground(lipgloss.Color("#960312"))
			}
			if rows[row-1][STATUS_COL] == HASH_MATCH {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("#43BF6D"))
			}
			if rows[row-1][STATUS_COL] == NEW_ENTRY {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("#4b13e8"))
			}
			return lipgloss.NewStyle().Foreground(lipgloss.Color("#4b13e8"))
		}).
		Headers(HEADER_FILE, HEADER_STATUS).
		Rows(rows...)

	fmt.Println(t)
}
