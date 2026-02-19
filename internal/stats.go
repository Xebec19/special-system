package internal

import (
	"fmt"
	"strings"
)

func stats(showLines, showWords, showBytes bool, files []string) string {

	var filesStat strings.Builder

	for _, val := range files {
		filesStat.WriteString(fmt.Sprintf("%s\n", val))
	}

	return filesStat.String()
}
