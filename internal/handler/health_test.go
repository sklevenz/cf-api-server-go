package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/handler"
)

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		accept       string
		wantBody     string
		wantMimeType string
	}{
		{"application/json", `{"status":"ok"}` + "\n", "application/json"},
		{"text/plain", "ok\n", "text/plain"},
		{"", "ok\n", "text/plain"},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		req.Header.Set("Accept", tc.accept)

		rr := httptest.NewRecorder()
		s := handler.Server{}
		s.GetHealthHandler(rr, req)

		res := rr.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		if ct := res.Header.Get("Content-Type"); !strings.HasPrefix(ct, tc.wantMimeType) {
			t.Errorf("unexpected content-type: got %s, want %s", ct, tc.wantMimeType)
		}
		if string(body) != tc.wantBody {
			t.Errorf("unexpected body for Accept=%s:\ngot  %q\nwant %q", tc.accept, string(body), tc.wantBody)
		}
	}
}
