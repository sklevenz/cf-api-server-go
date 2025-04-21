package integration_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

// Test for / (root) endpoint.
func TestVersionEndpoint(t *testing.T) {
	baseURL, shutdown := StartTestServer(t)
	defer shutdown()

	url, err := url.JoinPath(baseURL, "/version")
	if err != nil {
		t.Fatalf("failed to join URL path: %v", err)
	}

	resp, body := DoRequestWithResponse(t, url, map[string]string{
		httpx.HeaderAccept: httpx.ContentTypeJSON,
	})

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get(httpx.HeaderContentType); ct != httpx.ContentTypeJSON {
		t.Errorf("expected content type application/json, got %s", ct)
	}

	var jsonResponse map[string]string
	if err := json.Unmarshal([]byte(body), &jsonResponse); err != nil {
		t.Fatalf("failed to parse response body as JSON: %v", err)
	}

	if version, ok := jsonResponse["semanticVersion"]; !ok || version != "dev" {
		t.Errorf("expected JSON body to contain {\"semanticVersion\":\"dev\"}, got %s", body)
	}
}
