package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloudplay/auth"
	"cloudplay/session"
)

var (
	authAgent  auth.AuthenticationAgent
	sessionMgr session.Manager
)

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &creds)
		token, err := authAgent.Login(creds.Username, creds.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token)
	})

	mux.HandleFunc("/session/start", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth", http.StatusUnauthorized)
			return
		}
		ok, _ := authAgent.Verify(authHeader)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		var req struct {
			GameID string `json:"game_id"`
		}
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &req)
		s, err := sessionMgr.Start(req.GameID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
	})
}

func main() {
	authAgent = auth.NewSimpleAuth()
	sessionMgr = session.NewInMemoryManager()

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
