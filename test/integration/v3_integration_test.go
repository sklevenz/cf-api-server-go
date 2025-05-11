package integration_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

// Test for / (root) endpoint.
func TestV3RootEndpoint(t *testing.T) {
	baseURL, shutdown := StartTestServer(t)
	defer shutdown()

	u, _ := url.Parse(baseURL)
	u.Path = path.Join(u.Path, "v3")
	endpoint := u.String()

	resp, body := DoRequestWithResponse(t, endpoint, nil)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get(httpx.HeaderContentType); ct != httpx.ContentTypeJSON {
		t.Errorf("expected content type application/json, got %s", ct)
	}

	var jsonBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &jsonBody); err != nil {
		t.Errorf("expected valid JSON object, got: %v", body)
	}

	t.Logf("Body: %v", jsonBody)
}
