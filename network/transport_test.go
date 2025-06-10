package network

import "testing"

func TestDummyTransport(t *testing.T) {
	var tr Transport = DummyTransport{}
	if err := tr.Send([]byte("data")); err != nil {
		t.Fatal(err)
	}
}
