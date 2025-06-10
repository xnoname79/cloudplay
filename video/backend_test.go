package video

import (
	"io"
	"testing"
)

func TestDummyBackend(t *testing.T) {
	frames := [][]byte{[]byte("a"), []byte("b")}
	b := NewDummyBackend(frames)
	if err := b.Start(); err != nil {
		t.Fatal(err)
	}
	f1, err := b.ReadFrame()
	if err != nil || string(f1) != "a" {
		t.Fatalf("unexpected first frame %s %v", f1, err)
	}
	f2, err := b.ReadFrame()
	if err != nil || string(f2) != "b" {
		t.Fatalf("unexpected second frame %s %v", f2, err)
	}
	_, err = b.ReadFrame()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}
	if err := b.Stop(); err != nil {
		t.Fatal(err)
	}
}
