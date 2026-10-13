// Package hitdetection implements per-camera dart-impact detection:
// background subtraction to find changed regions within the calibrated
// board area, and a state machine that only reports a candidate hit once
// a blob has settled (filtering out hands, arms, and other transient
// motion). It assumes the board has already been located by
// targetdetection/calibration.
package hitdetection

import (
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// CameraDetector tracks hit-detection state for a single camera's frame
// stream.
type CameraDetector struct {
	camera config.CameraID
	cfg    config.HitDetectorConfig
	state  State
}

// NewCameraDetector builds a hit detector for one camera.
func NewCameraDetector(camera config.CameraID, cfg config.HitDetectorConfig) *CameraDetector {
	panic("not implemented")
}

// Feed processes the next frame from this camera, updating internal state.
// It returns a CandidateHit (ok == true) exactly when the state machine
// transitions into StateSettled.
func (d *CameraDetector) Feed(frame types.Frame) (hit types.CandidateHit, ok bool) {
	panic("not implemented")
}

// ResetBackground re-baselines the background model, e.g. after a
// confirmed hit or on a periodic timer, to absorb the dart now embedded
// in the board and lighting drift.
func (d *CameraDetector) ResetBackground(frame types.Frame) {
	panic("not implemented")
}

// State returns the detector's current state, mainly for debugging/UI.
func (d *CameraDetector) State() State {
	panic("not implemented")
}
