package resource

import "testing"

func TestDummyAllocator(t *testing.T) {
	var a Allocator = DummyAllocator{}
	if err := a.Check(); err != nil {
		t.Fatal(err)
	}
}
