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

func grep(pattern, text string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return re.MatchString(text)
}

func sed(old_pattern, new_pattern, line string) string {
	re := regexp.MustCompile(old_pattern)
	return re.ReplaceAllString(line, new_pattern)
}

func shellExec(command string) string {
	commandPieces := strings.Split(command, " ")
	result, err := exec.Command(commandPieces[0], commandPieces[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(string(result), "\n")
}

// byte256ToString: remove the blanks from string conversion
func byteToString(s string) string {
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

func getLogoDimensions(logo string) (width, height int) {
	width = 0
	height = 0
	for _, line := range strings.Split(logo, "\n") {
		line = lineStripColorCode(line)
		innerWidth := len(line)
		if innerWidth > width {
			width = innerWidth
		}
		height++
	}
	return width, height
}

func lineStripColorCode(line string) string {
	return sed(`\\033\[\d+m`, ``, line)
}

func getBoldTitle(title string) string {
	return fmt.Sprintf("%s%s%s%s", golorama.GetCSI(golorama.BOLD), golorama.GetCSI(golorama.YELLOW), title, reset())
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
