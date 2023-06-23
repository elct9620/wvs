package server

import (
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const SessionCookieName = "__WVSSSID"
const SessionExtendDuration = time.Hour * 1

type SessionStore interface {
	Renew(req *http.Request) *http.Cookie
	Find(id string) *Session
	Create(id, remoteAddr, userAgent string) *Session
	Destroy(id string) error
}

type InMemorySession struct {
	mu    sync.RWMutex
	items map[string]*Session
}

func NewInMemorySession() *InMemorySession {
	return &InMemorySession{
		items: make(map[string]*Session),
	}
}

func (s *InMemorySession) Find(id string) *Session {
	return s.items[id]
}

func (s *InMemorySession) Renew(req *http.Request) *http.Cookie {
	cookie, err := req.Cookie(SessionCookieName)
	if !errors.Is(err, http.ErrNoCookie) {
		return nil
	}

	newExpired := time.Now().Add(SessionExtendDuration)
	remoteAddr, _, _ := net.SplitHostPort(req.RemoteAddr)

	if IsValidSession(s, cookie, remoteAddr, req.UserAgent()) {
		return &http.Cookie{
			Name:     SessionCookieName,
			Value:    cookie.Value,
			HttpOnly: true,
			Expires:  newExpired,
		}
	}

	if cookie != nil {
		_ = s.Destroy(cookie.Value)
	}

	session := s.Create(uuid.NewString(), remoteAddr, req.UserAgent())

	return &http.Cookie{
		Name:     SessionCookieName,
		Value:    session.id,
		HttpOnly: true,
		Expires:  newExpired,
	}
}

func (s *InMemorySession) Create(id, remoteAddr, userAgent string) *Session {
	session := NewSession(id, remoteAddr, userAgent)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[id] = session
	return session
}

func (s *InMemorySession) Destroy(id string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if session, ok := s.items[id]; ok {
		if session.publishIO != nil {
			err = session.publishIO.Close()
		}
	}

	delete(s.items, id)
	return err
}

func IsValidSession(store SessionStore, cookie *http.Cookie, remoteAddr, userAgent string) bool {
	if cookie == nil {
		return false
	}

	if len(cookie.Value) == 0 {
		return false
	}

	session := store.Find(cookie.Value)
	if session == nil {
		return false
	}

	if !cmp.Equal(session.lastAddr, remoteAddr) {
		return false
	}

	if !cmp.Equal(session.userAgent, userAgent) {
		return false
	}

	return true
}
