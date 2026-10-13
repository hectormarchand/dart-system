// Package targetdetection locates the physical dartboard within a
// camera's frame and produces the pixel-space landmark points calibration
// needs to compute a homography. It runs once at startup and periodically
// thereafter to correct for camera drift, independently of hitdetection
// which assumes the board's position is already known.
//
// The board is located via ArUco markers glued around its frame: each
// marker is uniquely identified, and its 4 corners each become a
// Landmark, giving several well-distributed correspondence points per
// marker for calibration's homography computation.
//
// Detection goes through a small local cgo wrapper (aruco.cpp/aruco.h)
// instead of gocv.io/x/gocv/contrib: cgo builds every .cpp file in a
// package together, and that package also contains xfeatures2d.cpp, which
// needs a non-free contrib header Ubuntu's opencv package doesn't ship.
// The wrapper only touches opencv2/aruco.hpp, which is present.
package targetdetection

/*
#cgo pkg-config: opencv4
#cgo CXXFLAGS: --std=c++11
#include "aruco.h"
*/
import "C"

import (
	"fmt"
	"unsafe"

	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/types"
)

// arucoDictionary4x4_50 matches OpenCV's cv::aruco::DICT_4X4_50. 4x4
// markers stay reliably readable at the marker sizes suited to a
// dartboard-scale setup, and 50 unique IDs is far more than the few
// markers glued around one board.
const arucoDictionary4x4_50 = 0

// minMarkers is the minimum number of distinct ArUco markers that must be
// detected for a frame to be considered a confident board location. Below
// this, the resulting correspondence points are too few or too poorly
// distributed for a reliable homography.
const minMarkers = 3

// Landmark is a single recognizable point of the board (here, one corner
// of one ArUco marker) located in a camera's pixel space, tagged with the
// ID of its known counterpart in board space.
type Landmark struct {
	ID    string
	Point types.Point2D
}

// Detector locates the board within a single camera's frames.
type Detector struct {
	camera config.CameraID
}

// NewDetector builds a target detector for one camera.
func NewDetector(camera config.CameraID) *Detector {
	return &Detector{camera: camera}
}

// Detect attempts to locate the board in the given frame and returns its
// landmark points in pixel space. ok is false when the board could not be
// confidently located (bad lighting, obstruction, camera not yet aimed at
// the board, too few markers visible, ...).
func (d *Detector) Detect(frame types.Frame) (landmarks []Landmark, ok bool) {
	result := C.Aruco_DetectMarkers(C.Mat(unsafe.Pointer(frame.Mat.Ptr())), C.int(arucoDictionary4x4_50))
	defer C.ArucoMarkers_Close(result)

	count := int(result.count)
	if count < minMarkers {
		return nil, false
	}

	ids := unsafe.Slice(result.ids, count)
	corners := unsafe.Slice(result.corners, count*4*2)

	landmarks = make([]Landmark, 0, count*4)
	for i := 0; i < count; i++ {
		id := int(ids[i])
		for c := 0; c < 4; c++ {
			x := corners[i*8+c*2+0]
			y := corners[i*8+c*2+1]
			landmarks = append(landmarks, Landmark{
				ID:    fmt.Sprintf("marker-%d-corner-%d", id, c),
				Point: types.Point2D{X: float64(x), Y: float64(y)},
			})
		}
	}

	return landmarks, true
}
