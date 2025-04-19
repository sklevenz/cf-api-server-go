package server

import (
	"fmt"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/logger"
	middleware "github.com/sklevenz/cf-api-server/internal/middelware"
)

// NewHTTPServer creates and configures a new HTTP server.
func NewHTTPServer(addr string, cfgDir string) *http.Server {
	srv := handler.NewServer(cfgDir) // Initialize the API handler

	srv.LoadFavicon() // Load the favicon

	apiMux := http.NewServeMux()                        // Create a new HTTP request multiplexer
	apiHandler := generated.HandlerFromMux(srv, apiMux) // Bind generated routes to the mux

	outerMux := http.NewServeMux()
	outerMux.HandleFunc("/favicon.ico", srv.GetFaviconHandler) // custom route
	outerMux.HandleFunc("/health", srv.GetHealthHandler)       // custom route

	outerMux.Handle("/", apiHandler) // forward everything else to oapi

	wrappedHandler := middleware.LoggingMiddleware(outerMux) // Wrap the handler with logging middleware

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
