package input

import "testing"

func TestDummyProcessor(t *testing.T) {
	var p Processor = DummyProcessor{}
	if err := p.Handle([]byte("test")); err != nil {
		t.Fatal(err)
	}
}
