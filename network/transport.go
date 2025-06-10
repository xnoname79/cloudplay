package network

// Transport routes media packets.
type Transport interface {
	Send(packet []byte) error
}

type DummyTransport struct{}

func (d DummyTransport) Send(packet []byte) error {
	return nil
}
