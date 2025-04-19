package handler

import (
	"net/http"
	"strconv"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func (s *Server) GetFaviconHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// Respond to OPTIONS request
		w.Header().Set("Allow", "GET, OPTIONS")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if s.favicon == nil {
		http.NotFound(w, r)
		return
	}

	etag := httpx.GenerateETag(*s.favicon)
	if match := r.Header.Get("If-None-Match"); match == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", httpx.ContentTypeXIcon)
	w.Header().Set("Content-Length", strconv.Itoa(len(*s.favicon)))
	w.Header().Set("ETag", etag)
	w.WriteHeader(http.StatusOK)

	w.Write(*s.favicon)
}
