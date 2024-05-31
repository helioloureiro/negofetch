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

/****
# Flag:    --ascii_distro
# NOTE: AIX, Alpine, Anarchy, Android, Antergos, antiX, "AOSC OS",
#       "AOSC OS/Retro", Apricity, ArcoLinux, ArchBox, ARCHlabs,
#       ArchStrike, XFerience, ArchMerge, Arch, Artix, Arya, Bedrock,
#       Bitrig, BlackArch, BLAG, BlankOn, BlueLight, bonsai, BSD,
#       BunsenLabs, Calculate, Carbs, CentOS, Chakra, ChaletOS,
#       Chapeau, Chrom*, Cleanjaro, ClearOS, Clear_Linux, Clover,
#       Condres, Container_Linux, CRUX, Cucumber, Debian, Deepin,
#       DesaOS, Devuan, DracOS, DarkOs, DragonFly, Drauger, Elementary,
#       EndeavourOS, Endless, EuroLinux, Exherbo, Fedora, Feren, FreeBSD,
#       FreeMiNT, Frugalware, Funtoo, GalliumOS, Garuda, Gentoo, Pentoo,
#       gNewSense, GNOME, GNU, GoboLinux, Grombyang, Guix, Haiku, Huayra,
#       Hyperbola, janus, Kali, KaOS, KDE_neon, Kibojoe, Kogaion,
#       Korora, KSLinux, Kubuntu, LEDE, LFS, Linux_Lite,
#       LMDE, Lubuntu, Lunar, macos, Mageia, MagpieOS, Mandriva,
#       Manjaro, Maui, Mer, Minix, LinuxMint, MX_Linux, Namib,
#       Neptune, NetBSD, Netrunner, Nitrux, NixOS, Nurunner,
#       NuTyX, OBRevenge, OpenBSD, openEuler, OpenIndiana, openmamba,
#       OpenMandriva, OpenStage, OpenWrt, osmc, Oracle, OS Elbrus, PacBSD,
#       Parabola, Pardus, Parrot, Parsix, TrueOS, PCLinuxOS, Peppermint,
#       popos, Porteus, PostMarketOS, Proxmox, Puppy, PureOS, Qubes, Radix,
#       Raspbian, Reborn_OS, Redstar, Redcore, Redhat, Refracted_Devuan,
#       Regata, Rosa, sabotage, Sabayon, Sailfish, SalentOS, Scientific,
#       Septor, SereneLinux, SharkLinux, Siduction, Slackware, SliTaz,
#       SmartOS, Solus, Source_Mage, Sparky, Star, SteamOS, SunOS,
#       openSUSE_Leap, openSUSE_Tumbleweed, openSUSE, SwagArch, Tails,
#       Trisquel, Ubuntu-Budgie, Ubuntu-GNOME, Ubuntu-MATE, Ubuntu-Studio,
#       Ubuntu, Venom, Void, Obarun, windows10, Windows7, Xubuntu, Zorin,
#       and IRIX have ascii logos
# NOTE: Arch, Ubuntu, Redhat, and Dragonfly have 'old' logo variants.
#       Use '{distro name}_old' to use the old logos.
# NOTE: Ubuntu has flavor variants.
#       Change this to Lubuntu, Kubuntu, Xubuntu, Ubuntu-GNOME,
#       Ubuntu-Studio, Ubuntu-Mate  or Ubuntu-Budgie to use the flavors.
# NOTE: Arcolinux, Dragonfly, Fedora, Alpine, Arch, Ubuntu,
#       CRUX, Debian, Gentoo, FreeBSD, Mac, NixOS, OpenBSD, android,
#       Antrix, CentOS, Cleanjaro, ElementaryOS, GUIX, Hyperbola,
#       Manjaro, MXLinux, NetBSD, Parabola, POP_OS, PureOS,
#       Slackware, SunOS, LinuxLite, OpenSUSE, Raspbian,
#       postmarketOS, and Void have a smaller logo variant.
****/

const (
	// Horizontal spaces
	HSPACES int = 0
)

const (
	RESET = iota
	LIGHTRED
	LIGHTGREEN
	LIGHTYELLOW
	LIGHTBLUE
	LIGHTMAGENTA
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
)

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
	re, err := regexp.Compile(`\\033.*m[^m]`)
	if err != nil {
		log.Fatal("Failed to compile regex")
	}
	return re.ReplaceAllString(line, "")
}

func main() {
	dataFetch := negofetch{}
	dataFetch.detectOS()
	// logo := getLogo(dataFetch.OS)
	logo := macOSLogo()
	logoX, logoY := getLogoDimensions(logo)
	printLogo(logo)

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
	posX := logoX + HSPACES
	posY := termHeight - logoY
	setCursorPosition(posX, posY)
	positionStepUp(&posX, &posY)
	machineTag := getUsername() + "@" + getHostname()
	fmt.Printf("%s%s%s", golorama.GetCSI(golorama.LIGHTGREEN), machineTag, golorama.Reset())
	positionStepUp(&posX, &posY)

	for i := 0; i < len(machineTag); i++ {
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
	fmt.Printf("%s: %s", getBoldTitle("OS"), os)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Host"), host)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Kernel"), kernel)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Uptime"), uptime)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Packages"), packages)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Shell"), shell)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Resolution"), resolution)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("DE"), de)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("WM"), wm)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("WM Theme"), wmTheme)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Terminal"), terminal)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Terminal Font"), terminalFont)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("CPU"), cpu)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("GPU"), gpu)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Memory"), memory)

	// back to the end
	// fmt.Printf("\033[%d;%dH", termHeight, termWidth)
	setCursorPosition(termWidth, termHeight)
	// extra
}

func getBoldTitle(title string) string {
	return fmt.Sprintf("%s%s%s%s", golorama.GetCSI(golorama.BOLD), golorama.GetCSI(golorama.YELLOW), title, golorama.Reset())
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

	sysname := byteToString(string(uname.Sysname[:]))

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
