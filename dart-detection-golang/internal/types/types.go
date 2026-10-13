// Package types holds domain types shared across the detection pipeline
// packages (camera, detector, triangulation, hitsender).
package types

import (
	"time"

	"dart-detection-golang/internal/config"
	"gocv.io/x/gocv"
)

// Point2D is a pixel-space point in a single camera's frame.
type Point2D struct {
	X float64
	Y float64
}

// BoardPoint is a point in the shared board coordinate system, produced by
// triangulation, and matches the (x, y) contract expected by dart-backend's
// POST /api/dart-hit.
type BoardPoint struct {
	X float64
	Y float64
}

// CandidateHit is a single camera's report of a settled blob that might be
// a dart impact, before cross-camera triangulation.
type CandidateHit struct {
	Camera    config.CameraID
	Point     Point2D
	Timestamp time.Time
}

// DartHitEvent is the final, triangulated event describing a confirmed dart
// impact, ready to be sent to dart-backend.
type DartHitEvent struct {
	Position  BoardPoint
	Timestamp time.Time
}

// Frame wraps a single decoded camera frame with its source and capture
// time.
type Frame struct {
	Camera    config.CameraID
	Timestamp time.Time
	Mat       gocv.Mat
}
