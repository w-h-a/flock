// Package telemetry is the port interface for observability.
// Every scheduling decision, supervision transition, and health
// check is a wide event.
package telemetry

import "context"

// Telemetry emits wide events for scheduling and supervision.
type Telemetry interface {
	EmitSchedulingDecision(ctx context.Context, event SchedulingEvent) error
	EmitSupervisionTransition(ctx context.Context, event SupervisionEvent) error
}

type SchedulingEvent struct {
	Workload string
	Node     string
	Reason   string // "rendezvous_hash", "excess_drain", "node_failure"
}

type SupervisionEvent struct {
	Workload string
	Node     string
	OldState string
	NewState string
	ExitCode int
	Restarts uint
}
