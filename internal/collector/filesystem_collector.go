package collector

import (
	"context"
	"io/fs"
	"path/filepath"

	"ForLinux/internal/model"
)

// FileSystemCollector walks a filesystem and records file metadata
type FileSystemCollector struct {
	Root string
}

func (c *FileSystemCollector) Name() string {
	return "filesystem"
}

func (c *FileSystemCollector) Collect(ctx context.Context) ([]model.Event, error) {
	var events []model.Event

	filepath.Walk(c.Root, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		events = append(events, model.Event{
			Timestamp: info.ModTime(),
			Type:      model.EventFileSystem,
			Action:    "file_seen",
			Path:      path,
		})
		return nil
	})

	return events, nil
}
