package linux

import (
	"log"
	"os"
	"strings"

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
