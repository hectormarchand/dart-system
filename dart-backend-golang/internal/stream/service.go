// Package stream holds the latest JPEG frame received from each of the three
// board cameras and re-serves them as MJPEG streams to the frontend.
package stream

import "sync"

type Camera string

const (
	Left  Camera = "left"
	Front Camera = "front"
	Right Camera = "right"
)

type Service struct {
	mu     sync.RWMutex
	frames map[Camera][]byte
}

func NewService() *Service {
	return &Service{
		frames: map[Camera][]byte{
			Left:  {},
			Front: {},
			Right: {},
		},
	}
}

func (s *Service) SaveCameraFrames(left, front, right []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.frames[Left] = left
	s.frames[Front] = front
	s.frames[Right] = right
}

func (s *Service) GetFrame(camera Camera) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.frames[camera]
}
