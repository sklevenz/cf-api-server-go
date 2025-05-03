package server

import (
	"fmt"
	"net/http"

	gen "github.com/sklevenz/cf-api-server/internal/gen"
	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/logger"
	middleware "github.com/sklevenz/cf-api-server/internal/middelware"
)

// NewHTTPServer creates and configures a new HTTP server.
func NewHTTPServer(addr string, cfgDir string, semver string) (*http.Server, error) {
	srv, err := handler.NewServer(cfgDir, semver) // Initialize the API handler

	if err != nil {
		return nil, err
	}

	apiMux := http.NewServeMux()                  // Create a new HTTP request multiplexer
	apiHandler := gen.HandlerFromMux(srv, apiMux) // Bind generated routes to the mux

	outerMux := http.NewServeMux()
	outerMux.HandleFunc("/favicon.ico", srv.GetFaviconHandler) // custom route
	outerMux.HandleFunc("/health", srv.GetHealthHandler)       // custom route
	outerMux.HandleFunc("/version", srv.GetVersionHandler)     // custom route

	outerMux.Handle("/", apiHandler) // forward everything else to oapi

	// 404-Fallback-Wrapper
	fallbackWrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, pattern := outerMux.Handler(r)
		if pattern == "" || (r.URL.Path != "/" && pattern == "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})

	wrappedHandler := middleware.LoggingMiddleware(fallbackWrapper) // Wrap the handler with logging middleware

	return &http.Server{
		Handler: wrappedHandler,
		Addr:    addr,
	}, nil
}

// StartHTTPServer starts the HTTP server and logs errors if it fails.
func StartHTTPServer(srv *http.Server) {
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Log.Error("Server failed: %v", err)
	}
}

// StartServer creates and starts the HTTP server on the given port.
func StartServer(port int, cfgDir string, semver string) error {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	server, err := NewHTTPServer(addr, cfgDir, semver)

	if err != nil {
		return err
	}

	StartHTTPServer(server)

	return nil
}
