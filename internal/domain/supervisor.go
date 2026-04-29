package domain

import "time"

// SupervisionState is the Bernstein state machine.
// Each allocation has one. Local to the node running it.
//
// References:
//   - djb's daemontools supervise(8)
//   - skarnet s6 supervision suite
type SupervisionState struct {
	State    ProcessState
	Restarts uint
	LastExit *ExitInfo
	Backoff  time.Duration
}

type ProcessState int

const (
	Starting ProcessState = iota
	Running
	Stopping
	Crashed
	Restarting
)

type ExitInfo struct {
	Code      int
	Signal    int
	Timestamp time.Time
}

// Transition computes the next supervision state given an event.
// Pure function — no side effects.
func Transition(current SupervisionState, event ProcessEvent) SupervisionState {
	return SupervisionState{} // TODO
}

type ProcessEvent int

const (
	Started ProcessEvent = iota
	Exited
	HealthPassed
	HealthFailed
	StopRequested
)
