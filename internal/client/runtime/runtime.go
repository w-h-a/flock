// Package runtime is the port interface for process execution.
// Implementations start/stop/health-check actual processes.
package runtime

import (
	"context"

	"github.com/w-h-a/flock/internal/domain"
)

// Runtime executes and supervises workload instances locally.
type Runtime interface {
	Start(ctx context.Context, alloc domain.Allocation, workload domain.Workload) error
	Stop(ctx context.Context, alloc domain.Allocation) error
	Running(ctx context.Context) []domain.Allocation
}
