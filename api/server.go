package api

import (
	"log"
	"net/http"
)

func NewServer() {
	mux := http.NewServeMux() // Create a ServeMux

	// Register handlers with the ServeMux
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/ascii-art", AsciiArtHandler)

	log.Println("Starting server on http://127.0.0.1:8088")

	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
