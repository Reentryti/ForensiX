package collector

import (
	"context"

	"ForLinux/internal/model"
	pkg "ForLinux/internal/collector/package"
)

// PackageCollector detects and collects events from available package managers
type PackageCollector struct{}

func (c *PackageCollector) Name() string {
	return "packages"
}

func (c *PackageCollector) Collect(ctx context.Context) ([]model.Event, error) {
	var events []model.Event

	managers := pkg.DetectManager()

	for _, mgr := range managers {
		var ev []model.Event
		var err error

		switch mgr {
		case pkg.APT:
			ev, err = pkg.CollectApt(ctx)
		case pkg.PACMAN:
			ev, err = pkg.CollectPacman(ctx)
		case pkg.DNF:
			ev, err = pkg.CollectDnf(ctx)
		}

		if err == nil {
			events = append(events, ev...)
		}
	}

	return events, nil
}
