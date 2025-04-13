package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/pkg/httpx"
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

func TestGetRoot_ContentTypeJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	s := handler.Server{}
	s.GetRoot(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	contentType := resp.Header.Get(httpx.HeaderContentType)
	if contentType != httpx.ContentTypeJSON {
		t.Errorf("expected %q %q, got %q", httpx.HeaderContentType, httpx.ContentTypeJSON, contentType)
	}
}
