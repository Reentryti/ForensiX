package collector

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

type AptCollector struct {
}

func (c *AptCollector) Name() string {
	return "apt"
}

func (c *AptCollector) Collect(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open("/var/log/apt/history.log")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []model.Event
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Install") {
			events = append(events, model.Event{
				Timestamp: time.Now(),
				Type:      model.EventPackage,
				Action:    "package_install",
				Raw:       line,
			})
		}
	}
	return events, scanner.Err()
}
