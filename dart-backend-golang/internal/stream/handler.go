package stream

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"dart-backend-golang/internal/httpx"
)

const (
	routePrefix        = "/api/stream"
	mjpegFrameBoundary = "--frame"
	mjpegInterval      = 50 * time.Millisecond // ~10 fps

	maxUploadMemory = 32 << 20 // 32 MiB
)

func RegisterRoutes(mux *http.ServeMux, service *Service) {
	mux.HandleFunc("GET "+routePrefix+"/left-camera", handleMjpeg(service, Left))
	mux.HandleFunc("GET "+routePrefix+"/front-camera", handleMjpeg(service, Front))
	mux.HandleFunc("GET "+routePrefix+"/right-camera", handleMjpeg(service, Right))
	mux.HandleFunc("POST "+routePrefix+"/cameras", handleUploadCameraFrames(service))
}

func handleMjpeg(s *Service, camera Camera) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+mjpegFrameBoundary)
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "close")

		writeFrame := func() {
			buffer := s.GetFrame(camera)
			if len(buffer) == 0 {
				return
			}
			fmt.Fprintf(w, "%s\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\n\r\n", mjpegFrameBoundary, len(buffer))
			w.Write(buffer)
			w.Write([]byte("\r\n"))
			flusher.Flush()
		}

		writeFrame()

		ticker := time.NewTicker(mjpegInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				writeFrame()
			case <-r.Context().Done():
				return
			}
		}
	}
}

// handleUploadCameraFrames expects three JPEG frames uploaded as multipart
// form fields named "left", "front" and "right" (dart-detection is the only
// caller and is not wired up to this endpoint yet, so these field names are a
// deliberate, explicit contract rather than the original array-order one).
func handleUploadCameraFrames(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(maxUploadMemory); err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}

		left, err := readFormFile(r, string(Left))
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}
		front, err := readFormFile(r, string(Front))
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}
		right, err := readFormFile(r, string(Right))
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, err)
			return
		}

		s.SaveCameraFrames(left, front, right)

		httpx.JSON(w, map[string]string{"status": "ok"})
	}
}

func readFormFile(r *http.Request, field string) ([]byte, error) {
	file, _, err := r.FormFile(field)
	if err != nil {
		return nil, fmt.Errorf("missing or invalid form field %q: %w", field, err)
	}
	defer file.Close()

	return io.ReadAll(file)
}
