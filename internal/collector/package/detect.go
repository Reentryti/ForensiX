package pkg

import "os"

// Manager represents a package manager type
type Manager string

const (
	APT    Manager = "apt"
	DNF    Manager = "dnf"
	PACMAN Manager = "pacman"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DetectManager returns the list of package managers found on the system
func DetectManager() []Manager {
	var managers []Manager

	if exists("/var/log/apt/history.log") {
		managers = append(managers, APT)
	}
	if exists("/var/log/dnf.log") {
		managers = append(managers, DNF)
	}
	if exists("/var/log/pacman.log") {
		managers = append(managers, PACMAN)
	}
	return managers
}
