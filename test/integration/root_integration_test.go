package integration_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/sklevenz/cf-api-server/pkg/httpx"
	"github.com/sklevenz/cf-api-server/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelInfo, false, nil)
	os.Exit(m.Run())
}

// Test for / (root) endpoint.
func TestRootEndpoint(t *testing.T) {
	baseURL, shutdown := startTestServer(t)
	defer shutdown()

	resp, body := doRequestWithResponse(t, baseURL+"/")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get(httpx.HeaderContentType); ct != httpx.ContentTypeJSON {
		t.Errorf("expected content type application/json, got %s", ct)
	}

	var jsonBody map[string]interface{}
	if err := json.Unmarshal([]byte(body), &jsonBody); err != nil {
		t.Errorf("expected valid JSON object, got: %s", body)
	}
}
