package integration_test

import (
	"context"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/sklevenz/cf-api-server/internal/server"
)

const testCfgDir = "./test/cfg"

// startTestServer starts the server on a random port and returns the base URL and a shutdown function.
func startTestServer(t *testing.T) (baseURL string, shutdown func()) {
	// Listen on a random available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	// Create a new HTTP server using the actual listener address
	srv := server.NewHTTPServer(listener.Addr().String(), testCfgDir)

	// Run server in a goroutine
	go func() {
		_ = srv.Serve(listener)
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	return "http://" + listener.Addr().String(), func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}
}

// doRequest sends a GET request and returns the response body as a string.
func doRequestWithResponse(t *testing.T, url string) (*http.Response, string) {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	return resp, string(data)
}
