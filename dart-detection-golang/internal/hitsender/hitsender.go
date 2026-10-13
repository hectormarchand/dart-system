// Package hitsender delivers confirmed DartHitEvents to dart-backend over
// HTTP, decoupling the detection pipeline from the network I/O and its
// failure/retry handling.
package hitsender

import (
	"context"

	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// Sender posts dart hit events to dart-backend's POST /api/dart-hit.
type Sender struct {
	backendURL string
	client     any // *http.Client
}

// New builds a Sender targeting the configured backend URL.
func New(cfg config.Config) *Sender {
	panic("not implemented")
}

// Send delivers a single hit event, retrying transient failures as needed.
func (s *Sender) Send(ctx context.Context, event types.DartHitEvent) error {
	panic("not implemented")
}
