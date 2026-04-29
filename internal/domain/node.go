package domain

// Node represents a machine in the cluster.
// State derived from meld membership gossip.
type Node struct {
	ID        string
	Address   string
	Resources NodeResources
}

type NodeResources struct {
	CPUTotal      uint // MHz
	MemoryTotal   uint // MB
	CPUAvailable  uint
	MemoryAvailable uint
}
