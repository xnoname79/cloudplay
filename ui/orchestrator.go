package ui

// Orchestrator coordinates UI updates.
type Orchestrator interface {
	Notify(msg string) error
}

type DummyOrchestrator struct{}

func (d DummyOrchestrator) Notify(msg string) error { return nil }
