package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func (srv *Server) GetVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(srv.versionInfo)

}
