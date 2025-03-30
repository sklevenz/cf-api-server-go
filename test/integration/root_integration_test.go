package integration_test

import (
	"testing"
)

// Example test for /health endpoint.
func TestHealthEndpoint(t *testing.T) {
	baseURL, shutdown := startTestServer(t)
	defer shutdown()

	body := doRequest(t, baseURL+"/")
	if body != "{}\n" {
		t.Errorf("unexpected response: %s", body)
	}
}
