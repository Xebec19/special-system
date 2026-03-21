package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/Xebec19/special-system/pkg/logger"
)

func Stats(l, m, w, b bool, files []string) string {

	var filesStat strings.Builder

	showLines := l
	showChars := m
	showBytes := b
	showWords := w

	totalLines, totalChars, totalBytes := 0, 0, 0
	if !showLines && !showChars && !showBytes && !showWords {
		showLines = true
		showChars = true
		showBytes = true
	}

	for _, val := range files {
		// get lines, chars and size of given file
		lines, chars, words, size, err := stat(showLines, showChars, showWords, showBytes, val)
		if err != nil {
			logger.Error("error: failed to read file stats", err)
			fmt.Fprintln(&filesStat, "%w", err)
		}

		// sum total lines, chars and size
		totalLines += lines
		totalChars += chars
		totalBytes += size

		if showLines {
			filesStat.WriteString(strconv.FormatInt(int64(lines), 10))
			filesStat.WriteString(" ")
		}

		if showChars {
			filesStat.WriteString(strconv.FormatInt(int64(chars), 10))
			filesStat.WriteString(" ")
		}

		if showBytes {
			filesStat.WriteString(strconv.FormatInt(int64(size), 10))
			filesStat.WriteString(" ")
		}

		if showWords {
			filesStat.WriteString(strconv.FormatInt(int64(words), 10))
			filesStat.WriteString(" ")
		}

		filesStat.WriteString(val)
		filesStat.WriteString("\n")
	}

	if len(files) > 1 {
		fmt.Fprintln(&filesStat, totalLines, totalChars, totalBytes, "total")
	}

	return filesStat.String()
}

// stat returns no of lines, no of chars and size of a file
func stat(showLines, showChars, showWords, showBytes bool, fileName string) (int, int, int, int, error) {

	// dont do any processing if no param is required
	if !showLines && !showChars && !showBytes && !showWords {
		return 0, 0, 0, 0, nil
	}

	f, err := os.Open(fileName)
	if err != nil {
		logger.Log(fmt.Errorf("error: file can not be opened %w", err).Error())
		return 0, 0, 0, 0, err
	}

	defer f.Close()

	linesCount, charCount, wordCount, size := 0, 0, 0, 0
	inWord := false

	// get file size
	info, err := f.Stat()
	if err != nil {
		logger.Log("error: failed to get file stats ", err)
	} else if showBytes {
		size = int(info.Size()) // file size
	}

	reader := bufio.NewReader(f)

	for {
		// read each rune
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		// if file has atleast 1 char, it should have 1 lines as well
		if linesCount == 0 {
			linesCount++
		}

		if err != nil {
			return 0, 0, 0, 0, err
		}

		// count chars
		if showChars {
			charCount++
		}

		if showWords {
			if unicode.IsSpace(r) {
				inWord = false
			} else {
				if !inWord {
					wordCount++
					inWord = true
				}
			}
		}

		// count newlines
		if r == '\n' && showLines {
			linesCount++
		}
	}

	return linesCount, charCount, wordCount, size, nil
}
