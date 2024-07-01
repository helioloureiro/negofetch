package main

import (
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
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
