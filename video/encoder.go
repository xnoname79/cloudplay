package video

import "sync"

// Encoder encodes raw frames.
type Encoder interface {
	Encode(frame []byte) ([]byte, error)
}

type DummyEncoder struct{}

func (d DummyEncoder) Encode(frame []byte) ([]byte, error) {
	return frame, nil
}

// Streamer runs the encoding pipeline asynchronously. Frames submitted via
// Submit are encoded using the underlying Encoder and forwarded to the output
// channel returned by Out.
type Streamer struct {
	enc  Encoder
	in   chan []byte
	out  chan []byte
	stop chan struct{}
	wg   sync.WaitGroup
}

// NewStreamer creates a Streamer using the provided Encoder. The Streamer
// starts a goroutine that processes frames until Close is called.
func NewStreamer(enc Encoder) *Streamer {
	s := &Streamer{
		enc:  enc,
		in:   make(chan []byte, 16),
		out:  make(chan []byte, 16),
		stop: make(chan struct{}),
	}
	s.wg.Add(1)
	go s.loop()
	return s
}

// Submit queues a raw frame for encoding.
func (s *Streamer) Submit(frame []byte) {
	select {
	case s.in <- frame:
	case <-s.stop:
	}
}

// Out returns a read-only channel with encoded frames.
func (s *Streamer) Out() <-chan []byte { return s.out }

// Close stops the Streamer and waits for all queued frames to be processed.
func (s *Streamer) Close() {
	close(s.stop)
	s.wg.Wait()
	close(s.out)
}

func (s *Streamer) loop() {
	defer s.wg.Done()
	for {
		select {
		case frame := <-s.in:
			encoded, err := s.enc.Encode(frame)
			if err == nil {
				s.out <- encoded
			}
		case <-s.stop:
			return
		}
	}
}
