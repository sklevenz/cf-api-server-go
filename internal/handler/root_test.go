package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/handler"
)

func TestGetRoot(t *testing.T) {
	// Simulates a GET request to "/"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	s := handler.Server{}
	s.GetRoot(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Expects HTTP 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetRoot_Redirect(t *testing.T) {
	// Simulates a GET request with empty path
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.URL.Path = ""

	w := httptest.NewRecorder()

	s := handler.Server{}
	s.GetRoot(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Expects HTTP 301 Moved Permanently
	if resp.StatusCode != http.StatusMovedPermanently {
		t.Errorf("expected status 301, got %d", resp.StatusCode)
	}
	// Expects redirect location to "/"
	if loc := resp.Header.Get("Location"); loc != "/" {
		t.Errorf("expected redirect to '/', got %q", loc)
	}
}
