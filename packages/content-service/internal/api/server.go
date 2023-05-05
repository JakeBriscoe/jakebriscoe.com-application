package api

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Print("Setting server handlers")
	// Register HTTP request handlers
	http.HandleFunc("/tracks", FetchTracks)

	log.Print("Starting server")
	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
	log.Print("Server started on port 8080")
}
