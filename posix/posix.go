package posix

import (
	"log"
	"os"
	"os/user"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

func getTerminalSize() (int, int) {
	currentTerminalFD := int(os.Stdin.Fd())
	termWidth, termHeight, err := terminal.GetSize(currentTerminalFD)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return termWidth, termHeight
}

func getUptimeFromShell() string {
	uptime := shellExec("uptime")
	return strings.Split(uptime, ",")[0]
}

func getArchitecture() string {
	return shellExec("uname -m")
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

func (n *Negofetch) getShell() string {
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

	n.shell = shell
	return shell
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
