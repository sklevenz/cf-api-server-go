package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/handler"
)

// NewHTTPServer creates and configures a new HTTP server.
func NewHTTPServer(addr string) *http.Server {
	srv := handler.NewServer()                    // Initialize the API handler
	mux := http.NewServeMux()                     // Create a new HTTP request multiplexer
	handler := generated.HandlerFromMux(srv, mux) // Bind generated routes to the mux

	return &http.Server{
		Handler: handler,
		Addr:    addr,
	}
}

// StartHTTPServer starts the HTTP server and logs errors if it fails.
func StartHTTPServer(srv *http.Server) {
	log.Println("Starting server on", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

// StartServer creates and starts the HTTP server on the given port.
func StartServer(port int) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	server := NewHTTPServer(addr)
	StartHTTPServer(server)
}
