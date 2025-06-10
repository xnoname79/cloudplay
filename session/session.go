package session

import (
	"errors"
	"sync"
	"time"
)

// Session represents a running emulator session.
type Session struct {
	ID       string
	GameID   string
	Endpoint string
	Started  time.Time
}

type Manager interface {
	Start(gameID string) (*Session, error)
	Stop(id string) error
	Get(id string) (*Session, error)
}

type InMemoryManager struct {
	mu       sync.Mutex
	sessions map[string]*Session
}

func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{sessions: make(map[string]*Session)}
}

func (m *InMemoryManager) Start(gameID string) (*Session, error) {
	if gameID == "" {
		return nil, errors.New("game id required")
	}
	s := &Session{
		ID:       gameID + "_session",
		GameID:   gameID,
		Endpoint: "127.0.0.1:9000",
		Started:  time.Now(),
	}
	m.mu.Lock()
	m.sessions[s.ID] = s
	m.mu.Unlock()
	return s, nil
}

func (m *InMemoryManager) Stop(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.sessions[id]; !ok {
		return errors.New("session not found")
	}
	delete(m.sessions, id)
	return nil
}

func (m *InMemoryManager) Get(id string) (*Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[id]
	if !ok {
		return nil, errors.New("session not found")
	}
	return s, nil
}
