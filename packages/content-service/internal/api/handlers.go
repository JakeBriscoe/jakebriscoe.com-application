package api

import (
	"content-service/internal/database"
	"encoding/json"
	"fmt"
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s", r.URL.Path)
	fmt.Fprintf(w, "Hello, world!")
}
