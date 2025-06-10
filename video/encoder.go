package video

// Encoder encodes raw frames.
type Encoder interface {
	Encode(frame []byte) ([]byte, error)
}

type DummyEncoder struct{}

func (d DummyEncoder) Encode(frame []byte) ([]byte, error) {
	return frame, nil
}
