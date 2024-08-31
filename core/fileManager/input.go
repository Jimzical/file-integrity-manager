package fileManager

import (
	"bufio"
	"os"
	"strings"

	"github.com/Jimzical/file-integrity-manager/ui"
)

func Input(msg string) string {
	reader := bufio.NewReader(os.Stdin)

	ui.Highlight(msg)

	targetFolder, _ := reader.ReadString('\n')
	targetFolder = strings.TrimSpace(targetFolder)             // Remove all leading and trailing whitespace, including \r and \n
	targetFolder = strings.Replace(targetFolder, "\"", "", -1) // remove quotes
	return targetFolder
}
