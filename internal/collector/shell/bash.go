package shell

import (
	"bufio"
	"context"
	"os"
	"time"

	"ForLinux/internal/model"
)

// bashCollect parses a bash_history file into forensic events
func bashCollect(ctx context.Context, path string) ([]model.Event, error) {
	file, err := os.Open(path)
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

		cmd := scanner.Text()
		if cmd == "" {
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Now(),
			Type:      model.EventExecution,
			Action:    "shell_command",
			Source:    "bash",
			Command:  cmd,
			Raw:      cmd,
		})
	}
	return events, scanner.Err()
}
