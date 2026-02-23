package internal

import (
	"fmt"
	"strings"
)

func Stats(l, w, b bool, files []string) string {

	var filesStat strings.Builder

	showLines := l
	showChars := w
	showBytes := b

	if !showLines && !showChars && !showBytes {
		showLines = true
		showChars = true
		showBytes = true
	}

	for _, val := range files {
		fmt.Fprintf(&filesStat, "%t %t %t %s\n", showLines, showChars, showBytes, val)
	}

	if len(files) > 1 {
		filesStat.WriteString("total\n")
	}

	return filesStat.String()
}

func stat(showLines, showChars, showBytes bool, fileName string) (int, int, int) {


}
