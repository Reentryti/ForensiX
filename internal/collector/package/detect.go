package collector

import (
	"os"
)

// Passible package manager
type Manager string

const (
	APT    Manager = "apt"
	DNF    Manager = "dnf"
	PACMAN Manager = "pacman"
	//some others manager
)

// Check existence of the log file on system
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Get current package manager log file
// make an array cause some system can use multiple pkg manager
func DetectManager() []Manager {
	var managers []Manager

	if exists("") {
		managers = append(managers, APT)
	}
	if exists("") {
		managers = append(managers, DNF)
	}
	if exists("") {
		managers = append(managers, PACMAN)
	}
	return managers
}
