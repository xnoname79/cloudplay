package session

import "testing"

func TestInMemoryManager(t *testing.T) {
	m := NewInMemoryManager()
	s, err := m.Start("game1")
	if err != nil {
		t.Fatal(err)
	}
	got, err := m.Get(s.ID)
	if err != nil || got.GameID != "game1" {
		t.Fatalf("expected game1, got %v, err %v", got, err)
	}
	if err := m.Stop(s.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Get(s.ID); err == nil {
		t.Fatalf("expected error after stopping session")
	}
}
