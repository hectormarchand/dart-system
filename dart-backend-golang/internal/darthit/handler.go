package darthit

import (
	"encoding/json"
	"net/http"

	"dart-backend-golang/internal/httpx"
	"dart-backend-golang/internal/sse"
	"dart-backend-golang/internal/types"
)

const routePrefix = "/api/dart-hit"

func RegisterRoutes(mux *http.ServeMux, service *Service) {
	// Called from the opencv c++ module
	mux.HandleFunc("POST "+routePrefix, handleAddDartHit(service))
	mux.HandleFunc("GET "+routePrefix+"/stream-hits", handleStream(service))
}

func handleAddDartHit(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dartHitDto types.DartHitDto
		if err := json.NewDecoder(r.Body).Decode(&dartHitDto); err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}

		if _, err := s.AddDartHit(r.Context(), dartHitDto); err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}

		httpx.JSON(w, map[string]string{"status": "ok"})
	}
}

func handleStream(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sse.Serve(w, r, s.Broker(), "{}")
	}
}
