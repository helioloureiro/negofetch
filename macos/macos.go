package macos

import (
	"fmt"
	"strings"

	"github.com/fstanis/screenresolution"
	"github.com/helioloureiro/negofetch/posix"
	"github.com/helioloureiro/negofetch/utils"
)

func GetBrewPackages() string {
	packages := utils.ShellExec("brew list -1")
	counter := 0
	for _, pkg := range strings.Split(packages, "\n") {
		if utils.Grep("==>", pkg) {
			continue
		}
		if pkg != "" {
			counter++
		}
	}
	return fmt.Sprintf("%d (brew)", counter)

}

type MacOS struct{}

func (m *MacOS) GetUsername() string {
	return posix.GetUsername()
}

func (m *MacOS) GetHostname() string {
	return posix.GetHostname()
}

func (m *MacOS) GetShell() string {
	return posix.GetShell()
}

func (m *MacOS) GetMemory() string {
	return "memory"
}

func (m *MacOS) GetOS() string {
	return "os"
}

func (m *MacOS) GetUptime() string {
	return posix.GetUptimeFromShell()
}

func (m *MacOS) GetKernel() string {
	return posix.GetKernel()
}

func (m *MacOS) GetPackages() string {
	return "packages not implemented"
}

func (m *MacOS) GetScreenResolution() string {
	resolution := screenresolution.GetPrimary()
	return fmt.Sprintf("%dx%d", resolution.Width, resolution.Height)
}

/**************************************************
 * NOTE: I've not a macOS to keep testing for it  *
 * anymore, so this part of the code will be      *
 * abandoned.                                     *
 **************************************************/
