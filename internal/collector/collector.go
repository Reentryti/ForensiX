package collector

import (
	"context"

	"ForLinux/internal/model"
)

// Collector is the interface all forensic collectors must implement
type Collector interface {
	Name() string
	Collect(ctx context.Context) ([]model.Event, error)
}
