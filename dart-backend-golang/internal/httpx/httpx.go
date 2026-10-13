// Package httpx has small helpers shared by every HTTP handler.
package httpx

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error encoding JSON response:", err)
	}
}

func Error(w http.ResponseWriter, status int, err error) {
	log.Println(err)
	http.Error(w, err.Error(), status)
}
