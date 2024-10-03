package linux

import (
	"log"
	"os"
	"strings"
)

func getOSFromLSB() string {
	data, err := os.ReadFile("/etc/lsb-release")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if grep("DISTRIB_ID=", line) {
			system := strings.TrimPrefix(line, "DISTRIB_ID=")
			return sed(`"`, ``, system)
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
			os := strings.TrimPrefix(line, "NAME=")
			return sed(`"`, ``, os)
		}
	}
	return "Unknown Linux"
}
