package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gen "github.com/sklevenz/cf-api-server/internal/gen"
	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func TestGetV3Root(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v3", nil)
	w := httptest.NewRecorder()

	s, err := handler.NewTestServer()
	if err != nil {
		t.Fatalf("failed to create test server: %v", err)
	}
	s.GetApiRoot(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	contentType := resp.Header.Get(httpx.HeaderContentType)
	if contentType != httpx.ContentTypeJSON {
		t.Errorf("expected %q %q, got %q", httpx.HeaderContentType, httpx.ContentTypeJSON, contentType)
	}

	var root gen.Root
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if err := json.Unmarshal(body, &root); err != nil {
		t.Fatalf("failed to unmarshal response JSON: %v", err)
	}

	if root.Links.Self == nil {
		t.Fatal("expected self link to be present")
	}
	if !strings.HasPrefix(root.Links.Self.Href, "http://localhost") {
		t.Errorf("expected href to start with http://localhost, got %q", root.Links.Self.Href)
	}

	if root.Links.Self.Meta == nil {
		t.Fatal("expected meta to be present")
	}
}
