package ui

import "testing"

func TestDummyOrchestrator(t *testing.T) {
	var o Orchestrator = DummyOrchestrator{}
	if err := o.Notify("msg"); err != nil {
		t.Fatal(err)
	}
}
