package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sklevenz/cf-api-server/internal/httpx"
)

type Health struct {
	Status string `json:"status" xml:"status"`
}

func (Server) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{Status: "ok"}
	accept := r.Header.Get(httpx.HeaderAccept)

	switch accept {
	case "application/json":
		w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeJSON)
		json.NewEncoder(w).Encode(health)
	default:
		w.Header().Set(httpx.HeaderContentType, httpx.ContentTypeText)
		fmt.Fprintln(w, "ok")
	}
}
