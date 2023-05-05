package api

import (
	"content-service/internal/database"
	"encoding/json"
	"log"
	"net/http"
)

func FetchTracks(w http.ResponseWriter, r *http.Request) {
	log.Print("Processing request for FetchTracks")
	tracks, err := database.GetAllTracks()
	if err != nil {
		http.Error(w, "Failed to fetch tracks", http.StatusInternalServerError)
		return
	}

	// Return tracks as JSON
	jsonBytes, err := json.Marshal(tracks)
	if err != nil {
		http.Error(w, "Failed to encode tracks as JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
