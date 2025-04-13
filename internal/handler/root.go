package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/generated"
	"github.com/sklevenz/cf-api-server/pkg/httpx"
)

// GetRoot handles GET requests to the root endpoint ("/").
// If the request path is empty (r.URL.Path == ""), it redirects permanently to "/".
// Otherwise, it returns an empty JSON response with HTTP 200 OK.
func (Server) GetRoot(w http.ResponseWriter, r *http.Request) {
	var root generated.N200Root

	w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeJSON)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(root)
}
