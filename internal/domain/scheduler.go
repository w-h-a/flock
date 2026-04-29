package domain

// Schedule is a pure function. Given workloads and nodes, compute
// allocations via rendezvous hashing. Every node calling this with
// the same inputs gets the same outputs. No coordination needed.
//
// This is the core architectural bet: allocations are derived,
// not stored. Gossip converges the inputs (workloads, nodes).
// This function converges the outputs (allocations).
func Schedule(workloads []Workload, nodes []Node) []Allocation {
	return nil // TODO
}
