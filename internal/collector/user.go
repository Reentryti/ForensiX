package collector

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// UserCollector parses /etc/passwd for user enumeration
type UserCollector struct {
	Root string
}

func (c *UserCollector) Name() string {
	return "users"
}

func (c *UserCollector) Collect(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open(c.Root + "/etc/passwd")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var events []model.Event
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) < 7 {
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Now(),
			Type:      model.EventUser,
			Action:    "user_present",
			Result:    parts[0],
			Raw:       scanner.Text(),
		})
	}
	return events, scanner.Err()
}
