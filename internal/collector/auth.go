package collector

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// AuthLogCollector parses auth.log for authentication events
type AuthLogCollector struct {
	Path string
}

func NewAuthLogCollector(path string) *AuthLogCollector {
	return &AuthLogCollector{Path: path}
}

func (c *AuthLogCollector) Name() string {
	return "auth_log"
}

func (c *AuthLogCollector) Collect(ctx context.Context) ([]model.Event, error) {
	file, err := os.Open(c.Path)
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
		var severity model.Level

		switch {
		case strings.Contains(line, "Accepted"):
			action = "login_success"
			severity = model.LowLevel
		case strings.Contains(line, "Failed password"):
			action = "login_failure"
			severity = model.HighLevel
		case strings.Contains(line, "session opened"):
			action = "session_open"
			severity = model.LowLevel
		case strings.Contains(line, "session closed"):
			action = "session_close"
			severity = model.LowLevel
		default:
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Now(),
			Type:      model.EventAuth,
			Action:    action,
			Severity:  severity,
			Source:    "auth.log",
			Raw:       line,
		})
	}
	return events, scanner.Err()
}
