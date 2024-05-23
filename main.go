package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/helioloureiro/golorama"
	"golang.org/x/crypto/ssh/terminal"
)

/*
 7806     ¦   "mac"* | "Darwin")
 7807     ¦   ¦   set_colors 2 3 1 1 5 4
 7808     ¦   ¦   read -rd '' ascii_data <<'EOF'
 7809 ${c1}                    'c.
 7810     ¦   ¦   ¦   ¦,xNMM.
 7811     ¦   ¦   ¦  .OMMMMo
 7812     ¦   ¦   ¦  OMMM0,
 7813     ¦.;loddo:' loolloddol;.
 7814    cKMMMMMMMMMMNWMMMMMMMMMM0:
 7815 ${c2} .KMMMMMMMMMMMMMMMMMMMMMMMWd.
 7816  XMMMMMMMMMMMMMMMMMMMMMMMX.
 7817 ${c3};MMMMMMMMMMMMMMMMMMMMMMMM:
 7818 :MMMMMMMMMMMMMMMMMMMMMMMM:
 7819 ${c4}.MMMMMMMMMMMMMMMMMMMMMMMMX.
 7820  kMMMMMMMMMMMMMMMMMMMMMMMMWd.
 7821  ${c5}.XMMMMMMMMMMMMMMMMMMMMMMMMMMk
 7822   .XMMMMMMMMMMMMMMMMMMMMMMMMK.
 7823     ${c6}kMMMMMMMMMMMMMMMMMMMMMMd
 7824     ¦;KMMMMMMMWXXWMMMMMMMk.
 7825     ¦  .cooc,.    .,coo:.
 7826 EOF
*/

const (
	// Horizontal spaces
	HSPACES int = 5
)

func setColors(colors ...int) []string {
	result := make([]string, len(colors)+1)
	for id, colorID := range colors {
		result[id+1] = golorama.GetCSI(90 + colorID)
	}
	return result
}

func macOSLogo() string {
	return golorama.GetCSI(golorama.LIGHTGREEN_EX) + `
                    'c.
                  ,xNMM.
                .OMMMMo
                OMMM0,
      .;loddo:' loolloddol;.
    cKMMMMMMMMMMNWMMMMMMMMMM0:
` + golorama.GetCSI(golorama.LIGHTYELLOW_EX) + `  .KMMMMMMMMMMMMMMMMMMMMMMMWd.
  XMMMMMMMMMMMMMMMMMMMMMMMX.
` + golorama.GetCSI(golorama.LIGHTRED_EX) + ` ;MMMMMMMMMMMMMMMMMMMMMMMM:
 :MMMMMMMMMMMMMMMMMMMMMMMM:
 .MMMMMMMMMMMMMMMMMMMMMMMMX.
  kMMMMMMMMMMMMMMMMMMMMMMMMWd.
` + golorama.GetCSI(golorama.LIGHTMAGENTA_EX) + `  .XMMMMMMMMMMMMMMMMMMMMMMMMMMk
   .XMMMMMMMMMMMMMMMMMMMMMMMMK.
` + golorama.GetCSI(golorama.LIGHTBLUE_EX) + `     kMMMMMMMMMMMMMMMMMMMMMMd
      ;KMMMMMMMWXXWMMMMMMMk.
        .cooc,.    .,coo:.
 ` + golorama.Reset()
}

func otherMacOSLogo() string {
	// set_colors 2 3 1 1 5 4
	c := setColors(2, 3, 4, 1, 1, 5, 4)
	return `
 ` + c[1] + `                    'c.
                  ,xNMM.
                .OMMMMo
                OMMM0,
      .;loddo:' loolloddol;.
    cKMMMMMMMMMMNWMMMMMMMMMM0:
 ` + c[2] + ` .KMMMMMMMMMMMMMMMMMMMMMMMWd.
  XMMMMMMMMMMMMMMMMMMMMMMMX.
 ` + c[3] + `;MMMMMMMMMMMMMMMMMMMMMMMM:
 :MMMMMMMMMMMMMMMMMMMMMMMM:
 ` + c[4] + `.MMMMMMMMMMMMMMMMMMMMMMMMX.
  kMMMMMMMMMMMMMMMMMMMMMMMMWd.
  ` + c[5] + `.XMMMMMMMMMMMMMMMMMMMMMMMMMMk
   .XMMMMMMMMMMMMMMMMMMMMMMMMK.
     ` + c[6] + `kMMMMMMMMMMMMMMMMMMMMMMd
      ;KMMMMMMMWXXWMMMMMMMk.
        .cooc,.    .,coo:.
` + golorama.Reset()

}

func printLogo(system string) {
	switch system {
	case "macOS":
		fmt.Println(macOSLogo())
	case "alternative-macOS":
		fmt.Println(otherMacOSLogo())
	default:
		fmt.Println("System not found")
	}

}

func getLogoDimensions(logo string) (length, height int) {
	// break the string per line and find the longes width
	length = 0
	height = 0
	for _, line := range strings.Split(logo, "\n") {
		line = lineStripColorCode(line)
		innerLength := len(line)
		fmt.Printf("line %d: %s (%d)\n", height, line, innerLength)
		if innerLength > length {
			length = innerLength
		}
		height++
	}
	return length, height
}

func lineStripColorCode(line string) string {
	re, err := regexp.Compile(`\\033.*m[^m]`)
	if err != nil {
		log.Fatal("Failed to compile regex")
	}
	return re.ReplaceAllString(line, "")
}

func main() {
	fmt.Println("vim-go")
	printLogo("macOS")
	printLogo("alternative-macOS")
	x, y := getLogoDimensions(otherMacOSLogo())
	fmt.Println("Width:", x)
	fmt.Println("length:", y)

	// time.Sleep(3 * time.Second)
	// fmt.Println("\033[25;17H")
	//fmt.Println("TESTING")
	//fmt.Println("\033[26;17H")

	/*
		src: https://stackoverflow.com/questions/33025599/move-the-cursor-in-a-c-program
				printf ( "\033[2J");//clear screen
			    printf ( "\033[25;1H");//move cursor to row 25 col 1
			    printf ( "OVW");
			    printf ( "\033[9;1H");//move cursor to row 9 col 1
			    printf ( "enter your text ");//prompt
			    //printf ( "%s", input);
			    printf ( "\033[9;17H");//move cursor to row 9 col 17
	*/
	currentTerminalFD := int(os.Stdin.Fd())
	termWidth, termHeight, err := terminal.GetSize(currentTerminalFD)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("terminal width:", termWidth)
	fmt.Println("terminal height:", termHeight)

	posX := x + HSPACES + 12
	posY := termHeight - y - 4
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	machineTag := "helio@machine"
	fmt.Printf("%s: %d", machineTag, termHeight-y-4)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++

	sizeTag := len(machineTag)
	for i := 0; i < sizeTag; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++

	os := "macOS"
	host := "machine"
	kernel := "Darwin 1.2.3"
	uptime := "5 days"
	packages := "10 (brew)"
	shell := "zsh 1.0"
	resolution := "1920x1080"
	de := "aqua"
	wm := "quartz"
	wmTheme := "graphite"
	terminal := "iterm2"
	terminalFont := "Monaco"
	cpu := "Apple M1"
	gpu := "Apple M1"
	memory := "16 GB"
	fmt.Printf("OS: %s", os)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Host: %s", host)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Kernel: %s", kernel)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Uptime: %s", uptime)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Packages: %s", packages)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Shell: %s", shell)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Resolution: %s", resolution)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("DE: %s", de)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("WM: %s", wm)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("WM Theme: %s", wmTheme)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Terminal: %s", terminal)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Terminal Font: %s\n", terminalFont)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("CPU: %s", cpu)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("GPU: %s", gpu)
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	fmt.Printf("Memory: %s", memory)

	// back to the end
	fmt.Printf("\033[%d;%dH", termHeight, termHeight)
}
