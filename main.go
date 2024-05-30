package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"

	// "github.com/capnm/sysinfo"
	"github.com/helioloureiro/golorama"
	//"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
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

const (
	RESET = iota
	LIGHTRED
	LIGHTGREEN
	LIGHTYELLOW
	LIGHTBLUE
	LIGHTMAGENTA
)

func setColors(colors ...int) []string {
	result := make([]string, len(colors)+1)
	for id, colorID := range colors {
		result[id+1] = colorConverter(colorID)
	}
	return result
}

func colorConverter(color int) string {
	switch color {
	case LIGHTRED:
		return golorama.GetCSI(golorama.LIGHTRED_EX)
	case LIGHTGREEN:
		return golorama.GetCSI(golorama.LIGHTGREEN_EX)
	case LIGHTYELLOW:
		return golorama.GetCSI(golorama.LIGHTYELLOW_EX)
	case LIGHTBLUE:
		return golorama.GetCSI(golorama.LIGHTBLUE_EX)
	case LIGHTMAGENTA:
		return golorama.GetCSI(golorama.LIGHTMAGENTA_EX)
	default:
		return golorama.GetCSI(golorama.RESET)
	}
}

func macOSLogo() string {
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
	default:
		fmt.Println("System not found:", system)
	}

}

func getLogoDimensions(logo string) (length, height int) {
	// break the string per line and find the longes width
	length = 0
	height = 0
	for _, line := range strings.Split(logo, "\n") {
		line = lineStripColorCode(line)
		innerLength := len(line)
		// fmt.Printf("line %d: %s (%d)\n", height, line, innerLength)
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
	x, y := getLogoDimensions(macOSLogo())

	dataFetch := negofetch{}
	dataFetch.detectOS()
	printLogo(dataFetch.OS)

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

	termWidth, termHeight := getTerminalSize()
	posX := x + HSPACES + 12
	posY := termHeight - y - 4
	fmt.Printf("\033[%d;%dH", posY, posX)
	posY++
	machineTag := getUsername() + "@" + getHostname()
	fmt.Printf("%s: %d", machineTag, termHeight-y-4)
	positionStepUp(&posX, &posY)

	sizeTag := len(machineTag)
	for i := 0; i < sizeTag; i++ {
		fmt.Printf("-")
	}
	positionStepUp(&posX, &posY)
	uname := getUname()
	os := uname.Sysname
	host := uname.Nodename
	kernel := fmt.Sprintf("%s %s", uname.Sysname, uname.Release)
	uptime := getUptime()
	packages := getPackages()
	shell := getShell()
	resolution := "1920x1080 (hardcoded)"
	de := "aqua (hardcoded)"
	wm := "quartz (hardcoded)"
	wmTheme := "graphite (hardcoded)"
	terminal := "iterm2 (hardcoded)"
	terminalFont := "Monaco (hardcoded)"
	cpu := "Apple M1 (hardcoded)"
	gpu := "Apple M1 (hardcoded)"
	memory := "16 GB (hardcoded)"
	fmt.Printf("%sOS%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), os)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sHost%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), host)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sKernel%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), kernel)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sUptime%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), uptime)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sPackages%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), packages)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sShell%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), shell)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sResolution%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), resolution)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sDE%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), de)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sWM%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), wm)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sWM Theme%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), wmTheme)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sTerminal%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), terminal)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sTerminal Font%s: %s\n", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), terminalFont)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sCPU%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), cpu)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sGPU%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), gpu)
	positionStepUp(&posX, &posY)
	fmt.Printf("%sMemory%s: %s", golorama.GetCSI(golorama.YELLOW), golorama.Reset(), memory)

	// back to the end
	// fmt.Printf("\033[%d;%dH", termHeight, termWidth)
	setCursorPosition(termWidth, termHeight)
	// extra
}

type negofetch struct {
	OS           string
	Host         string
	Kernel       string
	Uptime       string
	Packages     string
	Shell        string
	Resolution   string
	DE           string
	WM           string
	WMTheme      string
	Terminal     string
	TerminalFont string
	CPU          string
	GPU          string
	Memory       string
}

func (n *negofetch) detectOS() {
	if fileExist("/etc") {
		system := shellExec("uname -s")
		switch system {
		case "Darwin":
			n.OS = "macOS"
		case "Linux":
			if fileExist("/etc/lsb-release") {
				n.OS = getOSFromLSB()
			} else if fileExist("/etc/os-release") {
				n.OS = getOSFromOSRelease()
			}
		default:
			n.OS = "Unknown Unix: " + system
			fmt.Println("Sizeof: ", len(system))
		}

	} else {
		n.OS = "Windows (not found /etc)"
	}

}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func getOSFromLSB() string {
	data, err := os.ReadFile("/etc/lsb-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "DISTRIB_ID=") {
			return strings.TrimPrefix(line, "DISTRIB_ID=")
		}
	}
	return "Unknown"
}

func getOSFromOSRelease() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "NAME=") {
			return strings.TrimPrefix(line, "NAME=")
		}
	}
	return "Unknown"
}
func positionStepUp(x, y *int) {
	setCursorPosition(*x, *y)
	*y++
}

func setCursorPosition(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func getUsername() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return u.Username
}

func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return h
}

func getUname() unix.Utsname {
	var uname unix.Utsname
	err := unix.Uname(&uname)
	if err != nil {
		log.Fatal(err)
	}
	return uname
}

func getShell() string {
	shell := os.Getenv("SHELL")

	if grep("/bin/bash", shell) {
		shell = "bash"
	} else if grep("/bin/zsh", shell) {
		shell = "zsh"
	} else if grep("/bin/fish", shell) {
		shell = "fish"
	} else if grep("/bin/tcsh", shell) {
		shell = "tcsh"
	} else if grep("/bin/csh", shell) {
		shell = "csh"
	} else if grep("/bin/csh", shell) {
		shell = "csh"
	} else if grep("/bin/ksh", shell) {
		shell = "ksh"
	}

	return shell
}

func grep(pattern, text string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return re.MatchString(text)
}

func getUptime() string {
	uname := getUname()

	sysname := byte256ToString(uname.Sysname)

	switch sysname {
	case "Linux":
		// si := sysinfo.Get()
		// return si.Uptime.String()
		return "hardcoded uptime"

	case "Darwin":
		return getUptimeFromShell()
	default:
		return "uknown system: " + getUptimeFromShell()
	}

}

func getUptimeFromShell() string {
	uptime := shellExec("uptime")
	return strings.Split(uptime, ",")[0]
}

func shellExec(command string) string {
	commandPieces := strings.Split(command, " ")
	result, err := exec.Command(commandPieces[0], commandPieces[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(string(result), "\n")
}

func byte256ToString(b [256]byte) string {

	str := ""
	for i := 0; i < len(b); i++ {
		// remove padding 0s
		if b[i] == 0 {
			continue
		}
		str += string(b[i])
	}
	return str
}

func getPackages() string {
	packages := shellExec("brew list -1")
	counter := 0
	for _, pkg := range strings.Split(packages, "\n") {
		if grep("==>", pkg) {
			continue
		}
		if pkg != "" {
			counter++
		}
	}
	return fmt.Sprintf("%d (brew)", counter)
}

func getTerminalSize() (int, int) {
	currentTerminalFD := int(os.Stdin.Fd())
	termWidth, termHeight, err := terminal.GetSize(currentTerminalFD)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return termWidth, termHeight
}
