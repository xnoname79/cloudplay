package input

// Processor handles incoming input frames.
type Processor interface {
	Handle(frame []byte) error
}

type DummyProcessor struct{}

func (d DummyProcessor) Handle(frame []byte) error {
	// stub
	return nil
}
