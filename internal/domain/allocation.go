package domain

// Allocation is DERIVED, not stored or replicated.
// Computed locally by each node from converged workloads + nodes
// via the schedule() function. If inputs converge (CRDT guarantee),
// outputs converge (deterministic function).
type Allocation struct {
	Workload    string
	Instance    uint
	NodeID      string
	Supervision SupervisionState
}
