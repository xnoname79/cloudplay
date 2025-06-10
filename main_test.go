package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"cloudplay/auth"
	"cloudplay/session"
)

func TestHealthEndpoint(t *testing.T) {
	authAgent = auth.NewSimpleAuth()
	sessionMgr = session.NewInMemoryManager()

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

func TestLoginAndStartSession(t *testing.T) {
	authAgent = auth.NewSimpleAuth()
	sessionMgr = session.NewInMemoryManager()

	mux := http.NewServeMux()
	setupRoutes(mux)

	loginReq := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"u","password":"p"}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, loginReq)
	if w.Code != http.StatusOK {
		t.Fatalf("login failed: %d", w.Code)
	}
	token := w.Body.String()

	startReq := httptest.NewRequest(http.MethodPost, "/session/start", strings.NewReader(`{"game_id":"g"}`))
	startReq.Header.Set("Authorization", structFromResponse(token))
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, startReq)
	if w2.Code != http.StatusOK {
		t.Fatalf("start failed: %d", w2.Code)
	}
}

func structFromResponse(body string) string {
	var t auth.Token
	json.Unmarshal([]byte(body), &t)
	return t.AccessToken
}
