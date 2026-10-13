// Package streaming optionally republishes raw or annotated camera frames
// (e.g. for a live debug view or the frontend's calibration UI), kept
// separate from the detection hot path.
package streaming

import (
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// Server exposes live camera frames to external consumers.
type Server struct {
	cfg config.Config
}

// New builds a streaming Server from configuration.
func New(cfg config.Config) *Server {
	panic("not implemented")
}

// Publish makes a frame available to current subscribers.
func (s *Server) Publish(frame types.Frame) {
	panic("not implemented")
}

// Start begins serving the streaming endpoint(s).
func (s *Server) Start() error {
	panic("not implemented")
}

// Stop shuts the streaming server down.
func (s *Server) Stop() error {
	panic("not implemented")
}
