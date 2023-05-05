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
	log.Print("Sever possibly started")
	if err != nil {
		log.Panic("Server failed to start: ", err)
	}
	log.Print("Server started on port 8080")
}
