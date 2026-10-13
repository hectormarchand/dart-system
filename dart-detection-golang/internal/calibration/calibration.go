// Package calibration loads and applies the per-camera homography that
// maps pixel-space points onto the shared board coordinate system.
package calibration

import (
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/targetdetection"
	"dart-detection-golang/internal/types"
)

// Homography represents the planar transform from one camera's pixel
// space to board space.
type Homography struct {
	Camera config.CameraID
	Matrix [3][3]float64
}

// Store holds the loaded homography for every configured camera.
type Store struct {
	homographies map[config.CameraID]Homography
}

// Load reads calibration files (one per camera) from cfg.CalibrationDir.
func Load(cfg config.Config) (*Store, error) {
	panic("not implemented")
}

// Set stores (or replaces) the homography for its camera, e.g. right
// after a successful Calibrate call. Safe to call on a Store returned by
// Load with no prior homography for that camera yet.
func (s *Store) Set(h Homography) {
	if s.homographies == nil {
		s.homographies = make(map[config.CameraID]Homography)
	}
	s.homographies[h.Camera] = h
}

// Project maps a pixel-space point from the given camera onto board space.
func (s *Store) Project(camera config.CameraID, p types.Point2D) (types.BoardPoint, error) {
	panic("not implemented")
}

// Calibrate (re)computes and persists the homography for a single camera
// from the landmarks located by targetdetection, resolving each
// landmark's known board-space position via ReferencePoint. Called at
// startup and whenever targetdetection re-locates the board (drift
// correction), independent of hitdetection.
func Calibrate(camera config.CameraID, landmarks []targetdetection.Landmark) (Homography, error) {
	panic("not implemented")
}

// ReferencePoint returns the fixed, known board-space position of a given
// landmark ID (e.g. a ring intersection or fiducial marker), as defined by
// the board's physical geometry.
func ReferencePoint(landmarkID string) (types.BoardPoint, bool) {
	panic("not implemented")
}
