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

	apiMux := http.NewServeMux() // Create a new HTTP request multiplexer
	apiHandler := gen.HandlerWithOptions(srv, gen.StdHTTPServerOptions{
		BaseURL:    "/v3", // use "/api" if you want to mount the API under a prefix
		BaseRouter: apiMux,
	})

	outerMux := http.NewServeMux()
	outerMux.HandleFunc("/favicon.ico", srv.GetFaviconHandler) // custom non-OpenAPI route
	outerMux.HandleFunc("/health", srv.GetHealthHandler)       // custom non-OpenAPI route
	outerMux.HandleFunc("/version", srv.GetVersionHandler)     // custom non-OpenAPI route

	outerMux.Handle("/v3/", apiHandler)      // forward all other requests to OpenAPI handler
	outerMux.HandleFunc("/", srv.GetApiRoot) //

	// 404 fallback wrapper
	fallbackWrapper := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h, pattern := outerMux.Handler(r)
		if pattern == "" || (r.URL.Path != "/" && pattern == "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})

	wrappedHandler := middleware.LoggingMiddleware(fallbackWrapper) // add logging middleware around the final handler

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
