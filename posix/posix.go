package posix

import (
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/helioloureiro/negofetch/utils"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

// GetTerminalSize: it returns the current terminal size
func GetTerminalSize() (int, int) {
	currentTerminalFD := int(os.Stdin.Fd())
	width, height, err := terminal.GetSize(currentTerminalFD)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return width, height
}

// GetUptimeFromShell: it uses command `uptime` and return the result
func GetUptimeFromShell() string {
	uptime := utils.ShellExec("uptime -p")
	return strings.Split(uptime, ",")[0]
}

// GetArchitecture: it returns the result from `uname -m`
func GetArchitecture() string {
	return utils.ShellExec("uname -m")
}

// GetUsername: it returns the username from whom is running the program
func GetUsername() string {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return u.Username
}

// GetHostname: it returns the hostname
func GetHostname() string {
	h, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return h
}

// GetUname: it return the `uname` on Go!
func GetUname() unix.Utsname {
	var uname unix.Utsname
	err := unix.Uname(&uname)
	if err != nil {
		log.Fatal(err)
	}
	return uname
}

// GetShell: it returns the current shell
func GetShell() string {
	shell := os.Getenv("SHELL")

	if utils.Grep("/bin/bash", shell) {
		shell = "bash"
	} else if utils.Grep("/bin/zsh", shell) {
		shell = "zsh"
	} else if utils.Grep("/bin/fish", shell) {
		shell = "fish"
	} else if utils.Grep("/bin/tcsh", shell) {
		shell = "tcsh"
	} else if utils.Grep("/bin/csh", shell) {
		shell = "csh"
	} else if utils.Grep("/bin/csh", shell) {
		shell = "csh"
	} else if utils.Grep("/bin/ksh", shell) {
		shell = "ksh"
	}

	return shell
}

// GetUptime: it gets uptime from system
func GetUptime() string {
	uname := GetUname()

	sysname := utils.ByteToString(string(uname.Sysname[:]))

	switch sysname {
	case "Linux":
		// si := sysinfo.Get()
		// return si.Uptime.String()
		return "hardcoded uptime"

	case "Darwin":
		return GetUptimeFromShell()
	default:
		return "uknown system: " + GetUptimeFromShell()
	}

}

// GetKernel: get the kernel version from uname
func GetKernel() string {
	return utils.ShellExec("uname -r")
}
