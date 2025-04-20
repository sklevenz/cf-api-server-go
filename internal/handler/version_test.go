package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func TestGetVersion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	s, err := handler.NewTestServer()
	if err != nil {
		t.Fatalf("failed to create test server: %v", err)
	}

	s.GetVersionHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get(httpx.HeaderContentType)
	if contentType != httpx.ContentTypeJSON {
		t.Errorf("expected %q %q, got %q", httpx.HeaderContentType, httpx.ContentTypeJSON, contentType)
	}
}
