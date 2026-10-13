package game

import (
	"encoding/json"
	"net/http"

	"dart-backend-golang/internal/httpx"
	"dart-backend-golang/internal/sse"
	"dart-backend-golang/internal/types"
)

const routePrefix = "/api/games"

func RegisterRoutes(mux *http.ServeMux, service *Service) {
	mux.HandleFunc("GET "+routePrefix, handleGetAll(service))
	mux.HandleFunc("POST "+routePrefix, handleCreate(service))
	mux.HandleFunc("GET "+routePrefix+"/current-active", handleGetActive(service))
	mux.HandleFunc("GET "+routePrefix+"/stream-current-active", handleStream(service))
	mux.HandleFunc("PATCH "+routePrefix+"/current-active", handlePatch(service))
}

func handleGetAll(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games, err := s.GetAllGames(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpx.JSON(w, games)
	}
}

func handleCreate(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto types.CreateDartGameDto
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}

		dartGame, err := s.CreateGame(r.Context(), dto)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpx.JSON(w, dartGame)
	}
}

func handleGetActive(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dartGame, err := s.GetActiveGame(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}
		httpx.JSON(w, dartGame)
	}
}

func handleStream(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dartGame, err := s.GetActiveGame(r.Context())
		initial := "{}"
		if err == nil && dartGame != nil {
			if data, err := json.Marshal(dartGame); err == nil {
				initial = string(data)
			}
		}
		sse.Serve(w, r, s.Broker(), initial)
	}
}

func handlePatch(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patchDto types.PatchGameDto
		if err := json.NewDecoder(r.Body).Decode(&patchDto); err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}

		if !isValidGamePatch(patchDto) {
			http.Error(w, "Invalid patch request", http.StatusBadRequest)
			return
		}

		currentActiveGame, err := s.GetActiveGame(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}
		if currentActiveGame == nil {
			http.Error(w, "No current active game", http.StatusBadRequest)
			return
		}

		patchedGame, err := s.PatchGame(r.Context(), currentActiveGame, patchDto)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, err)
			return
		}
		s.PushGameToClients(patchedGame)

		httpx.JSON(w, patchedGame)
	}
}

func isValidGamePatch(patch types.PatchGameDto) bool {
	return patch.Type == "next-player" || (patch.Type == "fix-score" && patch.Fix != nil)
}
