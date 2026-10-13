// Command dart-hit-faker stands in for dart-detection when no camera
// is available: it POSTs a random dart hit to dart-backend's /api/dart-hit
// on a fixed interval, so the rest of the system (scoring, SSE, frontend)
// can be exercised without physical hardware.
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	hitInterval    = 5 * time.Second
	requestTimeout = 3 * time.Second
)

// dartHitDto mirrors dart-backend's types.DartHitDto: r/theta describe
// the impact point in the board's polar coordinate system, estimatedPoints
// is the score dart-detection would have computed for that point.
type dartHitDto struct {
	R               float64 `json:"r"`
	Theta           float64 `json:"theta"`
	EstimatedPoints int     `json:"estimatedPoints"`
}

// possibleScores covers the values a real hit could produce: a miss, every
// single/double/triple segment, and the two bull values.
var possibleScores = buildPossibleScores()

func buildPossibleScores() []int {
	scores := []int{0, 25, 50}
	for segment := 1; segment <= 20; segment++ {
		scores = append(scores, segment, segment*2, segment*3)
	}
	return scores
}

func randomHit() dartHitDto {
	return dartHitDto{
		R:               rand.Float64() * 15,
		Theta:           rand.Float64() * 2 * math.Pi,
		EstimatedPoints: possibleScores[rand.Intn(len(possibleScores))],
	}
}

func main() {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://localhost:3000"
	}

	client := &http.Client{Timeout: requestTimeout}
	endpoint := backendURL + "/api/dart-hit"

	log.Printf("🎯 dart-hit-faker started, sending a fake hit to %s every %s", endpoint, hitInterval)

	ticker := time.NewTicker(hitInterval)
	defer ticker.Stop()

	for range ticker.C {
		sendFakeHit(client, endpoint)
	}
}

func sendFakeHit(client *http.Client, endpoint string) {
	hit := randomHit()

	body, err := json.Marshal(hit)
	if err != nil {
		log.Printf("failed to marshal fake hit: %v", err)
		return
	}

	resp, err := client.Post(endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Printf("failed to send fake hit: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.Printf("backend rejected fake hit (status %d): r=%.3f theta=%.3f points=%d", resp.StatusCode, hit.R, hit.Theta, hit.EstimatedPoints)
		return
	}

	log.Printf("sent fake hit: r=%.3f theta=%.3f points=%d", hit.R, hit.Theta, hit.EstimatedPoints)
}
