package monitoring

// Monitor collects metrics and logs.
type Monitor interface {
	Record(event string) error
}

type DummyMonitor struct{}

func (d DummyMonitor) Record(event string) error { return nil }
