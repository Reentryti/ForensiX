package pkg

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// CollectPacman parses pacman log for package events
func CollectPacman(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open("/var/log/pacman.log")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []model.Event
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return events, ctx.Err()
		default:
		}

		line := scanner.Text()

		if !strings.Contains(line, "[ALPM]") {
			continue
		}

		var action string
		switch {
		case strings.Contains(line, "installed"):
			action = "package_install"
		case strings.Contains(line, "upgraded"):
			action = "package_upgrade"
		case strings.Contains(line, "removed"):
			action = "package_remove"
		default:
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Now(),
			Type:      model.EventPackage,
			Action:    action,
			Source:    "pacman",
			Raw:       line,
		})
	}
	return events, scanner.Err()
}
