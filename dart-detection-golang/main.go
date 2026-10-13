package main

import (
	"dart-detection-golang/internal/config"
	"dart-detection-golang/internal/pipeline"
)

func main() {
	cfg := config.Load()

	p := pipeline.New(cfg)
	p.Run()
}
