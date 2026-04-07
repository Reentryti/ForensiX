package sysinfo

import (
	"errors"
	"os"
	"strings"

	"ForLinux/internal/model"
)

// collectDistro parses os-release files for distribution info
func (c *SystemInfoCollector) collectDistro() ([]model.Event, error) {
	paths := []string{
		c.Root + "/etc/os-release",
		c.Root + "/usr/lib/os-release",
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var events []model.Event
		lines := strings.Split(string(data), "\n")

		for _, l := range lines {
			if strings.Contains(l, "=") {
				parts := strings.SplitN(l, "=", 2)
				key := strings.TrimSpace(parts[0])
				val := strings.Trim(parts[1], `"'`)
				events = append(events, newEvent("distro."+key, val))
			}
		}
		return events, nil
	}
	return nil, errors.New("os-release not found")
}
