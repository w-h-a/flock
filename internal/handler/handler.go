// Package handler contains inbound adapters — the entry points
// for workload submissions and cluster queries.
package handler

import "context"

// Handler accepts workload commands (create, delete, list, status).
// Any node can handle any command — specs propagate via gossip.
type Handler interface {
	Create(ctx context.Context, spec []byte) error
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]byte, error)
	Status(ctx context.Context, name string) ([]byte, error)
}
