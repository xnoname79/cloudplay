package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// A small set of sample games that would be available for streaming.
var games = []string{
	"Gran Turismo 4",
	"Final Fantasy X",
	"Metal Gear Solid 3",
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// List available games
	mux.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string][]string{"games": games})
	})

	// Placeholder endpoint for starting a PS2 session
	mux.HandleFunc("/session/start", func(w http.ResponseWriter, r *http.Request) {
		game := r.URL.Query().Get("game")
		if game == "" {
			http.Error(w, "missing game parameter", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Starting PS2 session for %s... (not implemented)\n", game)
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
