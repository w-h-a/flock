// Package http is the inbound adapter for workload CRUD.
// Every flock node runs this handler. Accepts submissions,
// delegates to the reconciler's service methods.
package http

import "context"

// Handler serves workload operations over HTTP.
// Any node can handle any request. Specs propagate via gossip.
type Handler struct {
	// reconciler injected at construction
}

func (h *Handler) Create(ctx context.Context, spec []byte) error {
	return nil
}

func (h *Handler) Delete(ctx context.Context, name string) error {
	return nil
}

func (h *Handler) List(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (h *Handler) Status(ctx context.Context, name string) ([]byte, error) {
	return nil, nil
}
