package macos

import (
	"fmt"
	"strings"

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

/**************************************************
 * NOTE: I've not a macOS to keep testing for it  *
 * anymore, so this part of the code will be      *
 * abandoned.                                     *
 **************************************************/
