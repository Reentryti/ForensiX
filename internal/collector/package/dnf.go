package collector

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// DNF collector to collect package logs events from dnf logs
type DnfCollector struct {
}

func (c *DnfCollector) Name() string {
	return "dnf"
}

func (c *DnfCollector) Collect(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open("/var/log/dnf.log")
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
		var action string

		switch {
		case strings.Contains(line, "Installed:"):
			action = "package_install"
		case strings.Contains(line, "Upgrades:"):
			action = "package_upgrade"
		case strings.Contains(line, "Removed:"):
			action = "package_remove"
		default:
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Now(),
			Type:      model.EventPackage,
			Action:    action,
			Source:    "dnf",
			Raw:       line,
		})

	}
	return events, scanner.Err()
}
