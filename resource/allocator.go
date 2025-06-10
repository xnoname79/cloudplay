package resource

// Allocator scales resources based on metrics.
type Allocator interface {
	Check() error
}

type DummyAllocator struct{}

func (d DummyAllocator) Check() error { return nil }
