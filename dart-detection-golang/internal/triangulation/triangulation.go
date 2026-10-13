// Package triangulation aggregates settled CandidateHit reports coming
// from multiple cameras within a short time window into a single
// board-space DartHitEvent.
package triangulation

import (
	"dart-detection-golang/internal/calibration"
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// Aggregator buffers CandidateHit reports and groups the ones that belong
// to the same physical impact based on config.HitDetectorConfig.AggregationWindowMs.
type Aggregator struct {
	calibration *calibration.Store
	cfg         config.HitDetectorConfig

	pending []types.CandidateHit
}

// NewAggregator builds an Aggregator backed by the given calibration store.
func NewAggregator(store *calibration.Store, cfg config.HitDetectorConfig) *Aggregator {
	panic("not implemented")
}

// Add registers a new candidate hit from one camera. It returns a
// finalized DartHitEvent (ok == true) once enough corroborating cameras
// have reported within the aggregation window.
func (a *Aggregator) Add(candidate types.CandidateHit) (event types.DartHitEvent, ok bool) {
	panic("not implemented")
}

// Combine projects each candidate onto board space via calibration and
// merges them (e.g. averaging) into a single BoardPoint.
func Combine(store *calibration.Store, candidates []types.CandidateHit) (types.BoardPoint, error) {
	panic("not implemented")
}
