package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/helioloureiro/negofetch/linux"
	"github.com/helioloureiro/negofetch/macos"
	"github.com/helioloureiro/negofetch/posix"
	"github.com/helioloureiro/negofetch/utils"
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
	GREEN
	LIGHTYELLOW
	LIGHTBLUE
	LIGHTMAGENTA
	CYAN
	RED
	YELLOW
	BLUE
	MAGENTA
	LIGHTCYAN
	LIGHTGREEN
)

type Negofetch struct {
	username     string
	os           string
	hostname     string
	kernel       string
	uptime       string
	packages     string
	shell        string
	resolution   string
	deskEnviron  string
	wm           string
	wmTheme      string
	terminal     string
	terminalFont string
	cpu          string
	gpu          string
	memory       string
	screen       Screen
	logo         Logo
}

type Screen struct {
	leftSpace int // how much to move to the right to fit the logo
	x         int
	y         int
	width     int
	height    int
}

type Logo struct {
	name   string
	width  int
	height int
}

func (s *Screen) initialCursorPosition() {
	s.width, s.height = posix.GetTerminalSize()
}

func (s *Screen) moveCursorDown() {
	s.x++
}

var Version = "development"

func main() {
	distro := flag.String("ascii_distro", "", "ascii_distro")
	version := flag.Bool("version", false, "version")
	flag.Parse()

	if *version {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}

	s, r, m := cacheUname()
	fmt.Println("sizeof s:", len(s), s)
	fmt.Println("sizeof r:", len(r), r)
	fmt.Println("sizeof m:", len(m), m)

	negofetch := newNegoFetch()
	negofetch.screen.initialCursorPosition()
	negofetch.detectOS()
	if len(*distro) > 0 {
		fmt.Println("Distro:", *distro)
	}

	negofetch.printFormattedData()

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

	/*
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
		packages := dataFetch.getPackages()
		shell := dataFetch.getShell()
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
	*/
}

func newNegoFetch() Negofetch {
	system := NewOperatingSystem()
	return Negofetch{
		username: system.GetUsername(),
		hostname: system.GetHostname(),
		screen:   Screen{},
		logo:     Logo{},
		uptime:   system.GetUptime(),
		shell:    system.GetShell(),
		kernel:   system.GetKernel(),
		packages: system.GetPackages(),
	}
}

func cacheUname() (system, release, machine string) {
	// same as "uname -srm"
	u := posix.GetUname()
	return stringfy65(u.Sysname), stringfy65(u.Release), stringfy65(u.Machine)
}

func stringfy65(text [65]byte) string {
	result := ""
	for _, c := range text {
		if c != 0 {
			result = fmt.Sprintf("%s%s", result, string(c))
		}
	}
	return result
}

func (n *Negofetch) detectOS() {
	if utils.FileExist("/etc") {
		system := utils.ShellExec("uname -s")
		switch system {
		case "Darwin":
			n.os = "macOS"
		case "Linux":
			n.os = linux.GetDistro()
		default:
			n.os = "Unknown Unix: " + system
			fmt.Println("Sizeof: ", len(system))
		}

	} else {
		n.os = "Windows (not found /etc)"
	}

}

func positionStepUp(x, y *int) {
	setCursorPosition(*x, *y)
	*y++
}

func setCursorPosition(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func (n *Negofetch) getPackages() string {
	switch n.os {
	case "macOS":
		n.packages = macos.GetBrewPackages()
		return n.packages
	default:
		return "Not implemented yet for " + n.os

	}
}

type OperatingSystem interface {
	GetUsername() string
	GetHostname() string
	GetShell() string
	GetMemory() string
	GetOS() string
	GetUptime() string
	GetKernel() string
	GetPackages() string
}

func NewOperatingSystem() OperatingSystem {
	switch utils.ShellExec("uname -s") {
	case "Linux":
		return &linux.Linux{}
	case "macOS":
		return &macos.MacOS{}

	default:
		return &linux.Linux{}
	}
}

func (n *Negofetch) printFormattedData() {

	if n.logo.name == "" {
		n.screen.leftSpace = 1
		n.logo.width = 0
		n.logo.height = 0
	}

	/***********************************************************************
	helio@goosfraba
	---------------
	OS: Arch Linux x86_64
	Kernel: 6.11.1-arch1-1
	Uptime: 1 hour, 27 mins
	Packages: 2479 (pacman), 13 (flatpak)
	Shell: fish 3.7.1
	Resolution: 1920x1080, 1920x1080
	DE: Plasma 6.1.5
	WM: KWin
	WM Theme: Ambience-Dark
	Theme: Breeze Dark [Plasma], Breeze-Dark [GTK2], Breeze [GTK3]
	Icons: OpenDesktop-DarkIcons [Plasma], OpenDesktop-DarkIcons [GTK2/3]
	Terminal: yakuake
	CPU: AMD FX-8300 (8) @ 3.390GHz
	GPU: NVIDIA GeForce GTX 1050 Ti
	Memory: 8298MiB / 31998MiB
	************************************************************************/
	yPosition := returnStringOf(n.screen.leftSpace, " ")

	// title
	userData := fmt.Sprintf("%s@%s", n.username, n.hostname)
	userData = utils.GetBoldTitle(userData)
	fmt.Printf("%s%s\n", yPosition, userData)
	fmt.Printf("%s%s\n", yPosition, returnStringOf(len(userData), "-"))
	//---
	// OS
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("OS"), n.os)
	// Host
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Host"), n.hostname)
	// Kernel
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Kernel"), n.kernel)
	// Uptime
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Uptime"), n.uptime)
	// Packages
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Packages"), n.packages)
	// Shell
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Shell"), n.shell)
	// Resolution
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Resolution"), n.resolution)
	// Desktop Environment
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("DE"), n.deskEnviron)
	// Window Manager
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("WM"), n.wm)
	//
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Theme"), "none")
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Icons"), "none")
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Terminal"), "none")
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("CPU"), "none")
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("GPU"), "none")
	fmt.Printf("%s%s: %s\n", yPosition, utils.GetBoldTitle("Memory"), "none")

}

func returnStringOf(numberOfSpaces int, character string) string {
	word := character
	for i := numberOfSpaces; i > 0; i-- {
		word += character
	}
	return word
}
