package collector

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"ForLinux/internal/model"
)

func zshCollect(ctx context.Context, path string) ([]model.Event, error) {
	// Shell history file opening
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Line by line collection, store on slice
	var events []model.Event
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return events, ctx.Err()

		default:
		}

		// Querying the line to parse
		line := scanner.Text()

		// Check line format validation for zsh history
		if !strings.HasPrefix(line, ": ") {
			continue
		}

		// Line parsing (commande - metadata - timestamp)
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

		// Event creation (need it for the forensic)
		events = append(events, model.Event{
			Timestamp: time.Unix(ts, 0),
			Type:      model.EventExecution,
			Action:    "shell_command",
			Source:    "zsh",
			Command:   parts[1],
			Raw:       line,
		})
	}
	return events, scanner.Err()
}
