package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func TestGetRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	s, err := handler.NewTestServer()
	if err != nil {
		t.Fatalf("failed to create test server: %v", err)
	}
	s.GetRoot(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	contentType := resp.Header.Get(httpx.HeaderContentType)
	if contentType != httpx.ContentTypeJSON {
		t.Errorf("expected %q %q, got %q", httpx.HeaderContentType, httpx.ContentTypeJSON, contentType)
	}

	var root generated.N200Root
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if err := json.Unmarshal(body, &root); err != nil {
		t.Fatalf("failed to unmarshal response JSON: %v", err)
	}

	if root.Links == nil {
		t.Fatal("expected Links to be present")
	}
	if root.Links.Self == nil {
		t.Fatal("expected self link to be present")
	}
	if !strings.HasPrefix(root.Links.Self.Href, "http://localhost") {
		t.Errorf("expected href to start with http://localhost, got %q", root.Links.Self.Href)
	}

	if root.Links.Self.Method == nil || *root.Links.Self.Method != "GET" {
		t.Errorf("expected method to be GET, got %v", root.Links.Self.Method)
	}
	if root.Links.Self.Meta == nil {
		t.Fatal("expected meta to be present")
	}

	title, ok := (*root.Links.Self.Meta)["title"]
	if !ok || title != "Root endpoint" {
		t.Errorf("expected meta.title to be 'Root endpoint', got %v", title)
	}
}
