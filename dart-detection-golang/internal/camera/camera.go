// Package camera abstracts frame acquisition from the three physical
// cameras (or a static test image in dev mode) behind a common interface,
// and manages one capture goroutine per camera.
package camera

import (
	"context"

	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// Camera reads frames from a single physical (or simulated) video source.
type Camera interface {
	ID() config.CameraID
	Open() error
	Close() error
	// Read returns the next available frame, blocking until one is ready
	// or the context is cancelled.
	Read(ctx context.Context) (types.Frame, error)
}

// Manager owns the set of configured cameras and fans their frames out to
// consumers (detector, streaming).
type Manager struct {
	cameras map[config.CameraID]Camera
}

// NewManager constructs a Manager from the given configuration, opening a
// Camera implementation per configured device (or a test-image source in
// dev mode).
func NewManager(cfg config.Config) (*Manager, error) {
	panic("not implemented")
}

// Frames returns a channel of frames merged from all managed cameras.
func (m *Manager) Frames(ctx context.Context) <-chan types.Frame {
	panic("not implemented")
}

// Close releases all underlying camera resources.
func (m *Manager) Close() error {
	panic("not implemented")
}
