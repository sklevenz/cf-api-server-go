package server

import (
	"fmt"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/handler"
	middleware "github.com/sklevenz/cf-api-server/internal/middelware"
	"github.com/sklevenz/cf-api-server/pkg/logger"
)

// NewHTTPServer creates and configures a new HTTP server.
func NewHTTPServer(addr string, cfgDir string) *http.Server {
	srv := handler.NewServer(cfgDir)              // Initialize the API handler
	mux := http.NewServeMux()                     // Create a new HTTP request multiplexer
	handler := generated.HandlerFromMux(srv, mux) // Bind generated routes to the mux

	wrappedHandler := middleware.LoggingMiddleware(handler) // Wrap the handler with logging middleware

	return &http.Server{
		Handler: wrappedHandler,
		Addr:    addr,
	}
}

// StartHTTPServer starts the HTTP server and logs errors if it fails.
func StartHTTPServer(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Log.Error("Server failed: %v", err)
	}
}

// StartServer creates and starts the HTTP server on the given port.
func StartServer(port int, cfgDir string) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	server := NewHTTPServer(addr, cfgDir)
	StartHTTPServer(server)
}
