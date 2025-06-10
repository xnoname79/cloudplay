package video

import (
	"testing"
	"time"
)

func TestDummyEncoder(t *testing.T) {
	var e Encoder = DummyEncoder{}
	out, err := e.Encode([]byte("frame"))
	if err != nil || string(out) != "frame" {
		t.Fatalf("unexpected encode result: %s %v", out, err)
	}
}

func TestStreamer(t *testing.T) {
	s := NewStreamer(DummyEncoder{})
	defer s.Close()

	s.Submit([]byte("f"))

	select {
	case out := <-s.Out():
		if string(out) != "f" {
			t.Fatalf("unexpected output: %s", out)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for encoded frame")
	}
}
