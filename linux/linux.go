package linux

import (
	"log"
	"os"
	"strings"

	"github.com/helioloureiro/negofetch/utils"
)

func GetDistroFromLSB() string {
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

func GetDistroFromOSRelease() string {
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
