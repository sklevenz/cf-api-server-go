package integration_test

import (
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func TestFaviconEndpoint(t *testing.T) {
	baseURL, shutdown := StartTestServer(t)
	defer shutdown()

	url, err := url.JoinPath(baseURL, "/favicon.ico")
	if err != nil {
		t.Fatalf("failed to join URL path: %v", err)
	}

	resp, body := DoRequestWithResponse(t, url, nil)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	if ct := resp.Header.Get(httpx.HeaderContentType); ct != httpx.ContentTypeXIcon {
		t.Errorf("expected content type application/json, got %s", ct)
	}

	cl := resp.Header.Get(httpx.HeaderContentLength)
	if cl == "" {
		t.Errorf("expected content length header, got empty")
	}

	clInt, err := strconv.Atoi(cl)
	if err != nil {
		t.Errorf("failed to convert content length to integer: %v", err)
	}

	if len(body) != clInt {
		t.Errorf("size body (%v) != content length (%v)", len(body), clInt)
	}

}
