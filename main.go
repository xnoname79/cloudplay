package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Placeholder endpoint for starting a PS2 session
	mux.HandleFunc("/session/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Starting PS2 session... (not implemented)")
	})
}

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
