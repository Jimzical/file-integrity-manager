package fileManager

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	ui "github.com/Jimzical/file-integrity-manager/ui"
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
		Border(ui.TableBorderStyle).
		BorderStyle(ui.TableStyle).
		StyleFunc(func(row, col int) lipgloss.Style { // cant seem to remove this so the lipgloss dependency is still needed
			if row == 0 {
				// Style for the header row
				return ui.HeaderStyle
			}
			if rows[row-1][STATUS_COL] == HASH_MISMATCH {
				return ui.IncorrectStyle
			}
			if rows[row-1][STATUS_COL] == HASH_MATCH {
				return ui.SpecialStyle
			}
			if rows[row-1][STATUS_COL] == NEW_ENTRY {
				return ui.InfoStyle
			}
			return ui.InfoStyle
		}).
		Headers(HEADER_FILE, HEADER_STATUS).
		Rows(rows...)

	fmt.Println(t)
}
