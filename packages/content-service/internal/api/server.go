package api

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Print("Starting server")
	// Register HTTP request handlers
	http.HandleFunc("/tracks", FetchTracks)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	log.Print("Server started on port 8080")
}
