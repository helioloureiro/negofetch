package linux

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fstanis/screenresolution"
	"github.com/helioloureiro/negofetch/posix"
	"github.com/helioloureiro/negofetch/utils"
)

func GetDistro() string {
	if utils.FileExist("/etc/lsb-release") {
		return getDistroFromLSB()
	}

	if utils.FileExist("/etc/os-release") {
		return getDistroFromOSRelease()
	}
	return "Uknown distro"
}

func getDistroFromLSB() string {
	data, err := os.ReadFile("/etc/lsb-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if utils.Grep("DISTRIB_ID=", line) {
			system := strings.TrimPrefix(line, "DISTRIB_ID=")
			return utils.Sed(`"`, ``, system)
		}
	}
	return "Unknown"
}

func getDistroFromOSRelease() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "NAME=") {
			os := strings.TrimPrefix(line, "NAME=")
			return utils.Sed(`"`, ``, os)
		}
	}
	return "Unknown Linux"
}

type Linux struct{}

func (l *Linux) GetUsername() string {
	return posix.GetUsername()
}

func (l *Linux) GetHostname() string {
	return posix.GetHostname()
}

func (l *Linux) GetShell() string {
	return posix.GetShell()
}

func (l *Linux) GetMemory() string {
	return "mem"
}

func (l *Linux) GetOS() string {
	return "os"
}

func (l *Linux) GetUptime() string {
	return posix.GetUptimeFromShell()
}

func (l *Linux) GetKernel() string {
	return posix.GetKernel()
}

func (l *Linux) GetPackages() string {
	distro := GetDistro()
	distro = strings.ToLower(distro)

	var pkgSys string
	var pkgs []string

	switch distro {
	case "ubuntu":
		pkgSys = "apt/deb"
	case "debian":
		pkgSys = "apt/deb"
	case "mint":
		pkgSys = "apt/deb"
	case "suse":
		pkgSys = "zypper"
	case "fedora":
		pkgSys = "dnf"
	case "redhat":
		pkgSys = "dnf"
	case "archlinux":
		pkgSys = "pacman"
	default:
		pkgSys = "unknown"
	}
	pkgs = GetDistroPackages(distro)
	return fmt.Sprintf("%s (%d packages)", pkgSys, len(pkgs))

}

func GetDistroPackages(distro string) []string {
	var command string

	switch distro {
	case "ubuntu":
		return GetDebianPackages()
	case "debian":
		return GetDebianPackages()
	case "mint":
		return GetDebianPackages()
	case "archlinux":
		command = "pacman -Q"
		result := utils.ShellExec(command)
		return strings.Split(result, "\n")
	case "fedora":
		command = "rpm -qa"
		result := utils.ShellExec(command)
		return strings.Split(result, "\n")
	case "suse":
		command = "rpm -qa"
		result := utils.ShellExec(command)
		return strings.Split(result, "\n")
	case "unknown":
		return []string{}
	}

	// return empty
	return []string{}
}

func GetDebianPackages() []string {
	command := "dpkg -l"
	result := utils.ShellExec(command)
	var okLines []string
	for _, line := range strings.Split(result, "\n") {
		if utils.Grep(`^ii`, line) {
			okLines = append(okLines, line)
		}
	}
	return okLines
}

func (l *Linux) GetScreenResolution() string {
	resolution := screenresolution.GetPrimary()
	if resolution == nil {
		return ""
	}
	return fmt.Sprintf("%dx%d", resolution.Width, resolution.Height)
}
