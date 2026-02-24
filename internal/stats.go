package internal

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Xebec19/special-system/pkg/logger"
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

func stat(showLines, showChars, showBytes bool, fileName string) (int, int, int, error) {

	f, err := os.Open(fileName)
	if err != nil {
		logger.Log(fmt.Sprintf("error: file can not be opened %w", err))
		return -1, -1, -1, err
	}

	defer f.Close()

	buf := make([]byte, 4096) // 4kb buffer

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			logger.Log("error: failed to read file %s, %v", fileName, err)
			return -1, -1, -1, err
		}

		process(buf[:n])
	}

	return 0, 0, 0, nil
}

func process(data []byte) {
	logger.Log("Read bytes:", data)
}
