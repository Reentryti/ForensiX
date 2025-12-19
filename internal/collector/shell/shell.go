package collector

import (
	"context"
	"os"
	"path/filepath"

	"ForLinux/internal/model"
)

// Main collector for shell log like history collection
type ShellCollector struct {
	Home string
}

func NewShellCollector(home string) *ShellCollector {
	return &ShellCollector{Home: home}
}

func (c *ShellCollector) Name() string {
	return "shell_history"
}

// Verify working shell based on
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (c *ShellCollector) Collect(ctx context.Context) ([]model.Event, error) {
	//Event collect through
	var events []model.Event

	// In case of zsh shell usage
	zshPath := filepath.Join(c.Home, ".zsh_history")
	if exists(zshPath) {
		zshEvents, err := zshCollect(ctx, zshPath)
		if err == nil {
			events = append(events, zshEvents...)
		}
	}

	return events, nil

	// In case of bash shell usage
	//---implemented soon
}
