package integration_test

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/server"
	"github.com/sklevenz/cf-api-server/internal/testutil"
)

func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelDebug, false, nil)
	os.Exit(m.Run())
}

// startTestServer starts the server on a random port and returns the base URL and a shutdown function.
func StartTestServer(t *testing.T) (baseURL string, shutdown func()) {
	// Listen on a random available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	testCfgDir, err := testutil.GetTestDataPath("cfg")
	if err != nil {
		t.Fatalf("failed to get test data path: %v", err)
	}

	// Create a new HTTP server using the actual listener address
	srv, err := server.NewHTTPServer(listener.Addr().String(), testCfgDir, "dev")
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

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

func DoRequestWithResponse(t *testing.T, url string, headers map[string]string) (*http.Response, string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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
