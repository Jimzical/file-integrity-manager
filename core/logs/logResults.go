package logs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	ui "github.com/Jimzical/file-integrity-manager/ui"
	status "github.com/Jimzical/file-integrity-manager/core/status"
)

// getDisplayPath returns the display path of the file
//
// Parameters:
//   - filePath: The path of the file
//
// Returns:
//   - string: The display path of the file
//
// Example usage:
//
//     displayPath := getDisplayPath("/home/user/file.txt")
//     fmt.Println("Output: ",displayPath)
//     >> Output: .../user/file.txt
func GetDisplayPath(filePath string) string {
	components := strings.Split(filePath, string(filepath.Separator))
	if len(components) > 1 {
		lastFolder := components[len(components)-2]
		fileName := components[len(components)-1]
		return "..." + string(filepath.Separator) + lastFolder + string(filepath.Separator) + fileName
	}
	return filePath
}

func PrintTable(rows [][]string) {
	// if there are no rows, print a message and return
	if len(rows) == 0 {
		fmt.Println("No files to display")
		return
	}


	t := table.New().
		Border(ui.TableBorderStyle).
		BorderStyle(ui.TableStyle).
		StyleFunc(func(row, col int) lipgloss.Style { // cant seem to remove this so the lipgloss dependency is still needed
			if row == 0 {
				// Style for the header row
				return ui.HeaderStyle
			}

			chosenRow := rows[row-1]

			if chosenRow[STATUS_COL] == status.HASH_MISMATCH {
				return ui.IncorrectStyle
			}
			if chosenRow[STATUS_COL] == status.HASH_MATCH {
				return ui.SpecialStyle
			}
			if chosenRow[STATUS_COL] == status.NEW_ENTRY {
				return ui.InfoStyle
			}
			return ui.InfoStyle
		}).
		Headers(HEADER_FILE, HEADER_STATUS).
		Rows(rows...)

	fmt.Println(t)
}
