package auth

import (
	"errors"
	"sync"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type AuthenticationAgent interface {
	Login(username, password string) (*Token, error)
	Refresh(refreshToken string) (*Token, error)
	Verify(accessToken string) (bool, error)
}

type SimpleAuth struct {
	mu     sync.Mutex
	tokens map[string]*Token
}

func NewSimpleAuth() *SimpleAuth {
	return &SimpleAuth{tokens: make(map[string]*Token)}
}

func (s *SimpleAuth) Login(username, password string) (*Token, error) {
	if username == "" || password == "" {
		return nil, errors.New("missing credentials")
	}
	token := &Token{
		AccessToken:  username + "_access",
		RefreshToken: username + "_refresh",
		ExpiresAt:    time.Now().Add(time.Hour),
	}
	s.mu.Lock()
	s.tokens[token.AccessToken] = token
	s.mu.Unlock()
	return token, nil
}

func (s *SimpleAuth) Refresh(refreshToken string) (*Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, t := range s.tokens {
		if t.RefreshToken == refreshToken {
			delete(s.tokens, k)
			t.AccessToken = t.AccessToken + "_new"
			t.ExpiresAt = time.Now().Add(time.Hour)
			s.tokens[t.AccessToken] = t
			return t, nil
		}
	}
	return nil, errors.New("invalid refresh token")
}

func (s *SimpleAuth) Verify(accessToken string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.tokens[accessToken]
	if !ok {
		return false, nil
	}
	if time.Now().After(t.ExpiresAt) {
		delete(s.tokens, accessToken)
		return false, nil
	}
	return true, nil
}
