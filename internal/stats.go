package internal

import (
	"bytes"
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
		_, _, _, _ = stat(showLines, showChars, showBytes, val)
		fmt.Fprintf(&filesStat, "%t %t %t %s\n", showLines, showChars, showBytes, val)
	}

	if len(files) > 1 {
		filesStat.WriteString("total\n")
	}

	return filesStat.String()
}

// stat returns no of lines, no of chars and size of a file
func stat(showLines, showChars, showBytes bool, fileName string) (uint, uint, uint, error) {

	f, err := os.Open(fileName)
	if err != nil {
		logger.Log(fmt.Errorf("error: file can not be opened %w", err).Error())
		return 0, 0, 0, err
	}

	defer f.Close()

	buf := make([]byte, 4096) // 4kb buffer
	var linesCount uint = 0

	for {
		n, err := f.Read(buf)
		linesCount += uint(bytes.Count(buf[:n], []byte{'\n'}))
		if err == io.EOF {
			break
		}

		if err != nil {
			logger.Log("error: failed to read file %s, %v", fileName, err)
			return 0, 0, 0, err
		}

		process(buf[:n])
	}

	return linesCount, 0, 0, nil
}

func process(data []byte) {
	logger.Log("Read bytes:", string(data))
}
