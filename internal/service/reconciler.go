// Package service contains the controller/service layer.
// Thin glue: receives gossip updates, calls domain.Schedule(),
// diffs against running processes, starts/stops via runtime port.
package service

// Reconciler is the core loop:
//   1. Receive membership/workload changes via gossip
//   2. Call domain.Schedule() to compute desired allocations
//   3. Diff against runtime.Running() for actual allocations
//   4. Start missing, stop excess
//   5. Emit wide events
type Reconciler struct {
	// ports injected at construction
}
