package session_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elct9620/wvs/pkg/session"
)

func TestSessionNotSet(t *testing.T) {
	middleware := session.Middleware("1234567890123456")

	srv := httptest.NewServer(middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	found := false
	for _, cookie := range resp.Cookies() {
		if cookie.Name == session.DefaultCookieName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected to find cookie %s", session.DefaultCookieName)
	}
}

func TestSessionIsConfigured(t *testing.T) {
	middleware := session.Middleware("1234567890123456")

	srv := httptest.NewServer(middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Error(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  session.DefaultCookieName,
		Value: "vQbbUaO3tsF4fMlSyI7NvZH3ayZF2lEVlLrzPWrhE1W0/7+r8MxKluN6CVY+a/RHnpytQA==",
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	found := false
	for _, cookie := range resp.Cookies() {
		if cookie.Name == session.DefaultCookieName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected to find cookie %s", session.DefaultCookieName)
	}
}
