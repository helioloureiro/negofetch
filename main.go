package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	// "github.com/capnm/sysinfo"
	"github.com/helioloureiro/golorama"
	//"golang.org/x/crypto/ssh/terminal"
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

type Negofetch struct {
	OS           string
	OSType       string
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

func main() {
	distro := flag.String("ascii_distro", "", "ascii_distro")
	flag.Parse()
	negofetch := Negofetch{}
	negofetch.detectOS()
	var logo string
	if *distro != "" {
		fmt.Println("Using alternative logo: ", *distro)
		logo = getLogo(*distro)
	} else {
		fmt.Println("Using default logo: ", negofetch.OS)
		logo = getLogo(negofetch.OS)
	}
	logoX, logoY := getLogoDimensions(logo)
	uname := getUname()
	negofetch.OS = fmt.Sprintf("%s", uname.Sysname)
	negofetch.OSType = byteToString(fmt.Sprintf("%s", uname.Sysname))
	negofetch.Host = fmt.Sprintf("%s", uname.Nodename)
	negofetch.Kernel = fmt.Sprintf("%s %s", uname.Sysname, uname.Release)
	negofetch.Uptime = getUptime()
	negofetch.Packages = negofetch.getPackages()
	negofetch.Shell = negofetch.getShell()
	negofetch.Resolution = "1920x1080 (hardcoded)"
	negofetch.DE = "aqua (hardcoded)"
	negofetch.WM = "quartz (hardcoded)"
	negofetch.WMTheme = "graphite (hardcoded)"
	negofetch.Terminal = "iterm2 (hardcoded)"
	negofetch.TerminalFont = "Monaco (hardcoded)"
	negofetch.CPU = "Apple M1 (hardcoded)"
	negofetch.GPU = "Apple M1 (hardcoded)"
	negofetch.getMemory()

	endHere := true
	if endHere {
		fmt.Printf("%v", negofetch)
		return
	}
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
	fmt.Printf("%s: %s", getBoldTitle("OS"), negofetch.OS)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Host"), negofetch.Host)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Kernel"), negofetch.Kernel)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Uptime"), negofetch.Uptime)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Packages"), negofetch.Packages)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Shell"), negofetch.Shell)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Resolution"), negofetch.Resolution)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("DE"), negofetch.DE)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("WM"), negofetch.WM)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("WM Theme"), negofetch.WMTheme)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Terminal"), negofetch.Terminal)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Terminal Font"), negofetch.TerminalFont)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("CPU"), negofetch.CPU)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("GPU"), negofetch.GPU)
	positionStepUp(&posX, &posY)
	fmt.Printf("%s: %s", getBoldTitle("Memory"), negofetch.Memory)

	// back to the end
	// fmt.Printf("\033[%d;%dH", termHeight, termWidth)
	setCursorPosition(termWidth, termHeight)
	// extra
}

func (n *Negofetch) detectOS() {
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

func positionStepUp(x, y *int) {
	setCursorPosition(*x, *y)
	*y++
}

func setCursorPosition(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func (n *Negofetch) getPackages() string {
	switch n.OS {
	case "macOS":
		n.Packages = getBrewPackages()
		return n.Packages
	default:
		return "Not implemented yet for " + n.OS

	}
}

func (n *Negofetch) getMemory() {
	switch n.OSType {
	case "Linux":
		n.LinuxMemory()
	case "macOS":
		n.Memory = "not implemented, nor supported"
	default:
		n.Memory = "not implemented for " + n.OSType
	}
}

func (n *Negofetch) LinuxMemory() {
	fmt.Println("Getting memory on Linux")
	for _, line := range openFileAsArrayOfLines("/proc/meminfo") {
		if grep("MemTotal:", line) {
			// MemTotal:       24485228 kB
			line = sed("  *", " ", line)
			memTotalStr := strings.Split(line, " ")[1]
			memInt, err := strconv.Atoi(memTotalStr)
			if err != nil {
				panic(err)
			}
			var memUnit string = "KB"
			var sizeOfMem float64
			for _, unit := range []string{"MB", "GB", "TB", "PB"} {
				result := float64(memInt) / 1000.00
				fmt.Println(result)
				if result < 1000 {
					memUnit = unit
					sizeOfMem = result
					break
				}
				memInt = memInt / 1000
			}
			n.Memory = fmt.Sprintf("%0.2f %s", sizeOfMem, memUnit)
			break
		}
	}
}
