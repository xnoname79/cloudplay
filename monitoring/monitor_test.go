package monitoring

import "testing"

func TestDummyMonitor(t *testing.T) {
	var m Monitor = DummyMonitor{}
	if err := m.Record("event"); err != nil {
		t.Fatal(err)
	}
}
