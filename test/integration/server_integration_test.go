package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/sklevenz/cf-api-server/internal/server"
)

func TestServerIntegration(t *testing.T) {
	// Wait befor starting the server
	time.Sleep(1 * time.Second)

	// Start the server in a separate goroutine
	go func() {
		server.StartServer()
	}()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	// Make a GET request to the server
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	expectedBody := "Hello, World!\n"

	if string(body) != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, string(body))
	}
}
