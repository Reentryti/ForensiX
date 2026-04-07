package shell

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"ForLinux/internal/model"
)

// zshCollect parses a zsh_history file into forensic events
func zshCollect(ctx context.Context, path string) ([]model.Event, error) {
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

		line := scanner.Text()

		// Zsh extended history format: ": timestamp:0;command"
		if !strings.HasPrefix(line, ": ") {
			continue
		}

		parts := strings.SplitN(line[2:], ";", 2)
		if len(parts) != 2 {
			continue
		}

		meta := strings.Split(parts[0], ":")
		if len(meta) < 1 {
			continue
		}

		ts, err := strconv.ParseInt(meta[0], 10, 64)
		if err != nil {
			continue
		}

		events = append(events, model.Event{
			Timestamp: time.Unix(ts, 0),
			Type:      model.EventExecution,
			Action:    "shell_command",
			Source:    "zsh",
			Command:  parts[1],
			Raw:      line,
		})
	}
	return events, scanner.Err()
}
