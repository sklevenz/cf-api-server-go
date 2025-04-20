package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

// GetRootHandler handles GET requests to the root endpoint ("/").
// If the request path is empty (r.URL.Path == ""), it redirects permanently to "/".
// Otherwise, it returns an empty JSON response with HTTP 200 OK.
func (srv Server) GetRoot(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeJSON)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(srv.rootDocument)
}
