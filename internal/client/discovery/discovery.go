// Package discovery is the port interface for service discovery.
// The primary implementation serves DNS records derived from
// current allocations.
package discovery

import (
	"context"

	"github.com/w-h-a/flock/internal/domain"
)

// Discovery publishes workload locations so they can find each other.
type Discovery interface {
	Update(ctx context.Context, allocations []domain.Allocation, nodes []domain.Node) error
}
