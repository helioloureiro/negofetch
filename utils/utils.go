package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/helioloureiro/golorama"
)

// Grep: find whether a pattern exists in a string
func Grep(pattern, text string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return re.MatchString(text)
}

// Sed: replace a pattern in a string
func Sed(old_pattern, new_pattern, line string) string {
	re := regexp.MustCompile(old_pattern)
	return re.ReplaceAllString(line, new_pattern)
}

// ShellExec: execute some shell command and return the result
func ShellExec(command string) string {
	commandPieces := strings.Split(command, " ")
	result, err := exec.Command(commandPieces[0], commandPieces[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(string(result), "\n")
}

// Byte256ToString: remove the blanks from string conversion
func ByteToString(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		// remove padding 0s
		if s[i] == 0 {
			continue
		}
		str += string(s[i])
	}
	return str
}

// GetLogoDimension: return the width and height of a Logo not counting the terminal tags
func GetLogoDimensions(logo string) (width, height int) {
	width = 0
	height = 0
	for _, line := range strings.Split(logo, "\n") {
		line = LineStripColorCode(line)
		innerWidth := len(line)
		if innerWidth > width {
			width = innerWidth
		}
		height++
	}
	return width, height
}

// LineStripColorCode: remove the terminal tags from a string - NOT WORKING
func LineStripColorCode(line string) string {
	return Sed(`\\033\[\d+m`, ``, line)
}

// GetBoldTitle: title in bold style
func GetBoldTitle(title string) string {
	return fmt.Sprintf("%s%s%s%s", golorama.GetCSI(golorama.BOLD), golorama.GetCSI(golorama.YELLOW), title, golorama.GetCSI(golorama.RESET_ALL))
}

// FileExist: check whether a file exists
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// OpenFileAsArrayOfLines: return file content as array of strings
func OpenFileAsArrayOfLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Failed to open: " + filename)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}
