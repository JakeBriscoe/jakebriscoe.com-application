package api

import (
	"log"
	"net/http"
)

func StartServer() {
	log.Print("Setting server handlers")
	// Register HTTP request handlers
	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/hello", HelloWorld)
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/tracks", FetchTracks)

	log.Print("Starting server")

	// Start Server
	log.Println("Starting Server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error running the server", err)
	}

	log.Print("Server started on port 8080")
}
