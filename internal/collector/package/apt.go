package pkg

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// CollectApt parses apt history log for package events
func CollectApt(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open("/var/log/apt/history.log")
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
		if strings.HasPrefix(line, "Install") {
			events = append(events, model.Event{
				Timestamp: time.Now(),
				Type:      model.EventPackage,
				Action:    "package_install",
				Source:    "apt",
				Raw:       line,
			})
		}
	}
	return events, scanner.Err()
}
