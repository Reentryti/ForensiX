package shell

import (
	"context"
	"os"
	"path/filepath"

	"ForLinux/internal/model"
)

// ShellCollector collects shell history from bash and zsh
type ShellCollector struct {
	Home string
}

func NewShellCollector(home string) *ShellCollector {
	return &ShellCollector{Home: home}
}

func (c *ShellCollector) Name() string {
	return "shell_history"
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (c *ShellCollector) Collect(ctx context.Context) ([]model.Event, error) {
	var events []model.Event

	// Collect zsh history
	zshPath := filepath.Join(c.Home, ".zsh_history")
	if exists(zshPath) {
		zshEvents, err := zshCollect(ctx, zshPath)
		if err == nil {
			events = append(events, zshEvents...)
		}
	}

	// Collect bash history
	bashPath := filepath.Join(c.Home, ".bash_history")
	if exists(bashPath) {
		bashEvents, err := bashCollect(ctx, bashPath)
		if err == nil {
			events = append(events, bashEvents...)
		}
	}

	return events, nil
}
