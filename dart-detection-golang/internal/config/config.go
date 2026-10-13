// Package config centralizes environment-driven configuration for the
// detection pipeline (backend URL, camera device paths, dev mode, tuning
// thresholds for the detector state machine).
package config

// CameraID identifies one of the fixed camera positions around the board.
type CameraID string

const (
	CameraLeft  CameraID = "left"
	CameraFront CameraID = "front"
	CameraRight CameraID = "right"
)

// Config holds all runtime configuration for the detection pipeline.
type Config struct {
	// Dev enables local dev mode (e.g. reading from a static test image
	// instead of live camera devices, debug windows, verbose logging).
	Dev bool

	// BackendURL is the dart-backend base URL hits are POSTed to.
	BackendURL string

	// CameraDevices maps each camera to its video device path.
	CameraDevices map[CameraID]string

	// CalibrationDir is where per-camera homography files are stored.
	CalibrationDir string

	// Hit detector tuning.
	HitDetector HitDetectorConfig
}

// HitDetectorConfig holds the tunable thresholds for the per-camera
// motion/settle state machine used by hitdetection.
type HitDetectorConfig struct {
	// MinBlobArea / MaxBlobArea bound the plausible size of a dart impact.
	MinBlobArea float64
	MaxBlobArea float64

	// SettleFrames is how many consecutive stable frames are required
	// before a candidate blob is considered a confirmed hit.
	SettleFrames int

	// AggregationWindowMs is the max time window across cameras within
	// which candidate hits are considered the same physical event.
	AggregationWindowMs int
}

// Load reads configuration from environment variables, falling back to
// sane defaults.
func Load() Config {
	panic("not implemented")
}
