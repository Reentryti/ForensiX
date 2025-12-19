package collector

import (
	"ForLinux/internal/model"
	"bufio"
	"os"
	"time"
)

// Bash history collector based on bash_history file
func bashCollect(ctx context.Context, path string) ([]model.Event, error) {
	// Bash history file opening
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Events collect by line
	var events []model.Event
	scanner := bufio.NewScanner(file)
	// line by line, in case of empty one pass
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

		// Event creation (for the forensic)
		events = append(events, model.Event{
			//actually no time need by get it somehow
			Timestamp: time.Now(),
			Type:      model.EventExecution,
			Action:    "shell_command",
			Source:    "bash",
			Command:   cmd,
			Raw:       line,
		})
	}
	return events, scanner.Err()
}
