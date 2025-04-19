package server_test

import (
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/logger"
	"github.com/sklevenz/cf-api-server/internal/server"
	"github.com/sklevenz/cf-api-server/internal/testutil"
)

func TestMain(m *testing.M) {
	logger.Log = logger.New(logger.LevelInfo, false, nil)
	os.Exit(m.Run())
}

func TestNewHTTPServer(t *testing.T) {
	// Define a test address
	addr := "127.0.0.1:1234"

	testCfgDir, err := testutil.GetTestDataPath("cfg")
	if err != nil {
		t.Fatalf("failed to get test data path: %v", err)
	}

	// Create a new HTTP server using the test address
	srv := server.NewHTTPServer(addr, testCfgDir)

	// Ensure the server is not nil
	if srv == nil {
		t.Fatal("Expected non-nil server")
	}

	// Verify that the server address is correctly set
	if srv.Addr != addr {
		t.Errorf("Expected address %s, got %s", addr, srv.Addr)
	}

	// Ensure the server has a valid HTTP handler
	if srv.Handler == nil {
		t.Error("Expected non-nil handler")
	}
}

func TestIntegrationServer(t *testing.T) {
	testCfgDir, err := testutil.GetTestDataPath("cfg")
	if err != nil {
		t.Fatalf("failed to get test data path: %v", err)
	}

	// Create a new HTTP server with a random available port
	srv := server.NewHTTPServer("127.0.0.1:0", testCfgDir)

	// Manually create a listener to retrieve the actual port
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close() // Ensure the listener is closed after the test

	// Start the server in a separate goroutine
	go srv.Serve(ln)
	defer srv.Close() // Gracefully shut down the server after the test

	// Build the full URL using the actual listener address
	url := "http://" + ln.Addr().String() + "/"

	// Send an HTTP GET request to the root path
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
