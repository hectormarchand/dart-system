// Package pipeline wires camera acquisition, target/hit detection,
// cross-camera triangulation, and hit delivery together, and owns the
// main run loop.
package pipeline

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"dart-detection-golang/internal/calibration"
	"dart-detection-golang/internal/camera"
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/hitdetection"
	"dart-detection-golang/internal/hitsender"
	"dart-detection-golang/internal/streaming"
	"dart-detection-golang/internal/targetdetection"
	"dart-detection-golang/internal/triangulation"
	"dart-detection-golang/internal/types"
)

// Pipeline is the top-level orchestrator: it reads frames from the camera
// Manager, locates the board via targetdetection to (re)compute
// calibration, feeds frames to per-camera hit detectors, aggregates
// settled candidates into confirmed hits via triangulation, and forwards
// those to the Sender. It also feeds raw frames to the streaming server.
type Pipeline struct {
	cfg config.Config

	cameras         *camera.Manager
	targetDetectors map[config.CameraID]*targetdetection.Detector
	calibration     *calibration.Store
	calibrated      map[config.CameraID]bool
	hitDetectors    map[config.CameraID]*hitdetection.CameraDetector
	aggregator      *triangulation.Aggregator
	sender          *hitsender.Sender
	streamServer    *streaming.Server
}

// New constructs a Pipeline from configuration, wiring up all
// sub-components. Fallible I/O (opening cameras, loading calibration
// files, starting the stream server) is deferred to Run.
func New(cfg config.Config) *Pipeline {
	p := &Pipeline{
		cfg:             cfg,
		targetDetectors: make(map[config.CameraID]*targetdetection.Detector, len(cfg.CameraDevices)),
		hitDetectors:    make(map[config.CameraID]*hitdetection.CameraDetector, len(cfg.CameraDevices)),
		calibrated:      make(map[config.CameraID]bool, len(cfg.CameraDevices)),
		sender:          hitsender.New(cfg),
		streamServer:    streaming.New(cfg),
	}

	for id := range cfg.CameraDevices {
		p.targetDetectors[id] = targetdetection.NewDetector(id)
		p.hitDetectors[id] = hitdetection.NewCameraDetector(id, cfg.HitDetector)
	}

	return p
}

// Run starts the capture/detect/aggregate/send loop and blocks until the
// process is signaled to stop.
func (p *Pipeline) Run() {
	cameras, err := camera.NewManager(p.cfg)
	if err != nil {
		log.Fatalf("pipeline: opening cameras: %v", err)
	}
	p.cameras = cameras

	calib, err := calibration.Load(p.cfg)
	if err != nil {
		log.Printf("pipeline: loading calibration: %v (starting uncalibrated, will recalibrate from detected markers)", err)
		calib = &calibration.Store{}
	}
	p.calibration = calib
	p.aggregator = triangulation.NewAggregator(calib, p.cfg.HitDetector)

	if err := p.streamServer.Start(); err != nil {
		log.Fatalf("pipeline: starting stream server: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	for frame := range p.cameras.Frames(ctx) {
		p.streamServer.Publish(frame)
		p.processFrame(ctx, frame)
	}

	if err := p.Close(); err != nil {
		log.Printf("pipeline: closing: %v", err)
	}
}

// processFrame routes a single frame to target detection (until its
// camera is calibrated) or hit detection (once it is), and forwards any
// resulting confirmed hit to the sender.
func (p *Pipeline) processFrame(ctx context.Context, frame types.Frame) {
	if !p.calibrated[frame.Camera] {
		p.tryCalibrate(frame)
		return
	}

	hd, ok := p.hitDetectors[frame.Camera]
	if !ok {
		return
	}

	candidate, settled := hd.Feed(frame)
	if !settled {
		return
	}
	hd.ResetBackground(frame)

	event, ready := p.aggregator.Add(candidate)
	if !ready {
		return
	}

	if err := p.sender.Send(ctx, event); err != nil {
		log.Printf("pipeline: sending hit for camera %s: %v", frame.Camera, err)
	}
}

// tryCalibrate attempts to locate the board in the given frame and, on
// success, computes and stores its camera's homography so that camera can
// switch over to hit detection.
func (p *Pipeline) tryCalibrate(frame types.Frame) {
	td, ok := p.targetDetectors[frame.Camera]
	if !ok {
		return
	}

	landmarks, located := td.Detect(frame)
	if !located {
		return
	}

	homography, err := calibration.Calibrate(frame.Camera, landmarks)
	if err != nil {
		log.Printf("pipeline: calibrating camera %s: %v", frame.Camera, err)
		return
	}

	p.calibration.Set(homography)
	p.calibrated[frame.Camera] = true
}

// Close releases all resources owned by the pipeline (cameras, streaming
// server, ...).
func (p *Pipeline) Close() error {
	var errs []error

	if p.streamServer != nil {
		if err := p.streamServer.Stop(); err != nil {
			errs = append(errs, err)
		}
	}
	if p.cameras != nil {
		if err := p.cameras.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
