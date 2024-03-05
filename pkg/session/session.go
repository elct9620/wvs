package session

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const DefaultCookieName = "SSID"

type sessionCtxKey struct{}

type OptionFn func(*Session)

type Session struct {
	cookieName    string
	encryptionKey string
}

func New(encryptionKey string, options ...OptionFn) *Session {
	session := &Session{
		cookieName:    DefaultCookieName,
		encryptionKey: encryptionKey,
	}

	for _, option := range options {
		option(session)
	}

	return session
}

func WithCookieName(name string) OptionFn {
	return func(s *Session) {
		s.cookieName = name
	}
}

func Get(ctx context.Context) string {
	id, ok := ctx.Value(sessionCtxKey{}).(string)
	if !ok {
		return ""
	}

	return id
}

func Middleware(encryptionKey string, options ...OptionFn) func(http.Handler) http.Handler {
	session := New(encryptionKey, options...)
	return session.Handler
}

func (s *Session) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Get(r.Context()) != "" {
			next.ServeHTTP(w, r)
			return
		}

		sid, encrypted, err := findOrCreateSessionId(r, s.encryptionKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  s.cookieName,
			Value: encrypted,
		})

		ctx := context.WithValue(r.Context(), sessionCtxKey{}, sid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func findOrCreateSessionId(r *http.Request, key string) (string, string, error) {
	cookie, err := r.Cookie(DefaultCookieName)
	if err != nil {
		return newSessionId(key)
	}

	return getSessionId(cookie.Value, key)
}

func newSessionId(key string) (string, string, error) {
	sid := uuid.NewString()
	encrypted, err := Encrypt([]byte(sid), []byte(key))
	if err != nil {
		return "", "", err
	}

	return sid, encrypted, nil
}

func getSessionId(encrypted string, key string) (string, string, error) {
	decrypted, err := Decrypt(encrypted, []byte(key))
	if err != nil {
		return "", "", err
	}

	return string(decrypted), encrypted, nil
}
