package sysinfo

import (
	"context"

	"ForLinux/internal/model"
)

// SystemInfoCollector gathers system metadata (distro, arch, partitions, etc.)
type SystemInfoCollector struct {
	Root string
}

func New(root string) *SystemInfoCollector {
	return &SystemInfoCollector{Root: root}
}

func (c *SystemInfoCollector) Name() string {
	return "system_info"
}

func (c *SystemInfoCollector) Collect(ctx context.Context) ([]model.Event, error) {
	var events []model.Event

	distroEvents, err := c.collectDistro()
	if err == nil {
		events = append(events, distroEvents...)
	}

	return events, nil
}

func newEvent(key, value string) model.Event {
	return model.Event{
		Type:   model.EventSystem,
		Action: key,
		Result: value,
	}
}
