package video

import (
	"io"
	"os/exec"
)

// Backend represents a running emulator that produces raw video frames.
type Backend interface {
	// Start launches the emulator process.
	Start() error
	// ReadFrame returns the next raw frame. io.EOF is returned when no frames remain.
	ReadFrame() ([]byte, error)
	// Stop terminates the emulator process.
	Stop() error
}

// DummyBackend is a simple in-memory implementation used for tests.
type DummyBackend struct {
	frames [][]byte
	idx    int
}

// NewDummyBackend creates a DummyBackend that will output the provided frames.
func NewDummyBackend(frames [][]byte) *DummyBackend {
	return &DummyBackend{frames: frames}
}

func (d *DummyBackend) Start() error { d.idx = 0; return nil }

func (d *DummyBackend) ReadFrame() ([]byte, error) {
	if d.idx >= len(d.frames) {
		return nil, io.EOF
	}
	f := d.frames[d.idx]
	d.idx++
	return f, nil
}

func (d *DummyBackend) Stop() error { return nil }

// PCSX2Backend launches a PCSX2 process. Frame capture integration is TBD.
type PCSX2Backend struct {
	exe string
	iso string
	cmd *exec.Cmd
}

// NewPCSX2Backend returns a new backend configured to run the given executable and ISO image.
func NewPCSX2Backend(exe, iso string) *PCSX2Backend {
	return &PCSX2Backend{exe: exe, iso: iso}
}

func (p *PCSX2Backend) Start() error {
	p.cmd = exec.Command(p.exe, "--fullscreen", "--nogui", "--iso", p.iso)
	return p.cmd.Start()
}

// ReadFrame currently returns io.EOF because frame capture is not yet implemented.
func (p *PCSX2Backend) ReadFrame() ([]byte, error) {
	return nil, io.EOF
}

func (p *PCSX2Backend) Stop() error {
	if p.cmd != nil && p.cmd.Process != nil {
		return p.cmd.Process.Kill()
	}
	return nil
}
