package logs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	status "github.com/Jimzical/file-integrity-manager/core/status"
	ui "github.com/Jimzical/file-integrity-manager/ui"
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
//	displayPath := getDisplayPath("/home/user/file.txt")
//	fmt.Println("Output: ",displayPath)
//	>> Output: .../user/file.txt
func GetDisplayPath(filePath string) string {
	components := strings.Split(filePath, string(filepath.Separator))
	if len(components) > 1 {
		lastFolder := components[len(components)-2]
		fileName := components[len(components)-1]
		return "..." + string(filepath.Separator) + lastFolder + string(filepath.Separator) + fileName
	}
	return filePath
}

func PrintTable(rows [][]string, statusType string) {
	fmt.Println()

	// if there are no rows, print a message and return
	if len(rows) == 0 {
		fmt.Print("No files to display in this table\n\n")
		return
	}

	tableValueStyle := func(row, col int) lipgloss.Style {
		if row == 0 {
			// Style for the header row
			return ui.HeaderStyle
		}

		chosenRow := rows[row-1]

		switch chosenRow[STATUS_COL] {
		case status.HASH_MISMATCH:
			return ui.DangerStyle
		case status.HASH_MATCH:
			return ui.SuccessStyle
		case status.NEW_ENTRY:
			return ui.InfoStyle
		default:
			return ui.InfoStyle
		}
	}

	t := table.New().
		Border(ui.TableBorderStyle).
		BorderStyle(ui.TableStyle).
		StyleFunc(tableValueStyle).
		Headers(FILE_HEADER, STATUS_HEADER).
		Rows(rows...)

	fmt.Println(t)
}
