// Package domain contains the pure functional core.
// No I/O, no infrastructure imports. Testable with plain inputs/outputs.
package domain

// Workload is an immutable spec describing what to run.
// Replicated across the cluster as an OR-Set CRDT via meld.
type Workload struct {
	Name      string
	Exec      string            // binary path or container image
	Count     uint              // desired instance count
	Resources Resources
	Health    HealthCheck
	Placement PlacementMode
	Meta      map[string]string
}

type Resources struct {
	CPU    uint // MHz
	Memory uint // MB
}

type HealthCheck struct {
	Type     HealthType
	Address  string
	Interval uint // seconds
}

type HealthType int

const (
	TCP HealthType = iota
	HTTP
)

type PlacementMode int

const (
	Available   PlacementMode = iota // AP — keep scheduling during partition
	Coordinated                      // CP — freeze during partition (V2)
)
