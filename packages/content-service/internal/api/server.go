package api

import (
	"net/http"
)

func StartServer() {
	// Register HTTP request handlers
	http.HandleFunc("/tracks", FetchTracks)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
