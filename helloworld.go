package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
        http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("Listening on localhost:%s", port)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("Hello world received a request.")
        target := os.Getenv("TARGET")
        if target == "" {
                target = "World! I sure hope you see this"
        }
        fmt.Fprintf(w, "Hello %s!\n", target)
}