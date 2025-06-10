package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	setupRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	expected := "{\"status\":\"ok\"}\n"
	if w.Body.String() != expected {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestGamesEndpoint(t *testing.T) {
	mux := http.NewServeMux()
	setupRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/games", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	expected := "{\"games\":[\"Gran Turismo 4\",\"Final Fantasy X\",\"Metal Gear Solid 3\"]}\n"
	if w.Body.String() != expected {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}

func TestSessionStart(t *testing.T) {
	mux := http.NewServeMux()
	setupRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/session/start?game=TestGame", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	expected := "Starting PS2 session for TestGame... (not implemented)\n"
	if w.Body.String() != expected {
		t.Fatalf("unexpected body: %s", w.Body.String())
	}
}
