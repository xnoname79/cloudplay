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

func TestLoginRefreshAndSessionLifecycle(t *testing.T) {
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
	startReq.Header.Set("Authorization", accessFromResponse(token))
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, startReq)
	if w2.Code != http.StatusOK {
		t.Fatalf("start failed: %d", w2.Code)
	}

	// refresh token
	var tkn auth.Token
	json.Unmarshal([]byte(token), &tkn)
	refreshReq := httptest.NewRequest(http.MethodPost, "/refresh", strings.NewReader(`{"refresh_token":"`+tkn.RefreshToken+`"}`))
	w3 := httptest.NewRecorder()
	mux.ServeHTTP(w3, refreshReq)
	if w3.Code != http.StatusOK {
		t.Fatalf("refresh failed: %d", w3.Code)
	}

	// get session
	var s session.Session
	json.Unmarshal(w2.Body.Bytes(), &s)
	getReq := httptest.NewRequest(http.MethodGet, "/session/get?id="+s.ID, nil)
	getReq.Header.Set("Authorization", accessFromResponse(w3.Body.String()))
	w4 := httptest.NewRecorder()
	mux.ServeHTTP(w4, getReq)
	if w4.Code != http.StatusOK {
		t.Fatalf("get failed: %d", w4.Code)
	}

	// stop session
	stopReq := httptest.NewRequest(http.MethodPost, "/session/stop", strings.NewReader(`{"session_id":"`+s.ID+`"}`))
	stopReq.Header.Set("Authorization", accessFromResponse(w3.Body.String()))
	w5 := httptest.NewRecorder()
	mux.ServeHTTP(w5, stopReq)
	if w5.Code != http.StatusOK {
		t.Fatalf("stop failed: %d", w5.Code)
	}
}

func accessFromResponse(body string) string {
	var t auth.Token
	json.Unmarshal([]byte(body), &t)
	return t.AccessToken
}
