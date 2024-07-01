package main

import (
	"fmt"
	"strings"
)

func getBrewPackages() string {
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
