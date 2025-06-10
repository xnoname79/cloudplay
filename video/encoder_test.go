package video

import "testing"

func TestDummyEncoder(t *testing.T) {
	var e Encoder = DummyEncoder{}
	out, err := e.Encode([]byte("frame"))
	if err != nil || string(out) != "frame" {
		t.Fatalf("unexpected encode result: %s %v", out, err)
	}
}
