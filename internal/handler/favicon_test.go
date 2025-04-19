package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/sklevenz/cf-api-server/internal/handler"
	"github.com/sklevenz/cf-api-server/internal/httpx"
)

func TestOptionsFavicon(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/favicon.ico", nil)
	w := httptest.NewRecorder()

	cfgDir := filepath.Join("..", "..", "testdata", "cfg")

	s := handler.NewServer(cfgDir)
	s.LoadFavicon()
	s.GetFaviconHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status 204 NoContent, got %d", resp.StatusCode)
	}
}
func TestPutFavicon(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/favicon.ico", nil)
	w := httptest.NewRecorder()

	cfgDir := filepath.Join("..", "..", "testdata", "cfg")

	s := handler.NewServer(cfgDir)
	s.LoadFavicon()
	s.GetFaviconHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status 204 NoContent, got %d", resp.StatusCode)
	}
}

func TestGetFaviconHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/favicon.ico", nil)
	w := httptest.NewRecorder()

	cfgDir := filepath.Join("..", "..", "testdata", "cfg")

	s := handler.NewServer(cfgDir)
	s.LoadFavicon()
	s.GetFaviconHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if len(body) == 0 {
		t.Errorf("expected non-empty response body")
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	etag := resp.Header.Get("ETag")
	if etag == "" {
		t.Error("expected ETag header to be set")
	}

	contentLengthStr := resp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		t.Fatal("Content-Length header is missing")
	}

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		t.Fatalf("invalid Content-Length: %v", err)
	}

	if contentLength != len(body) {
		t.Errorf("expected Content-Length %d, got %d", len(body), contentLength)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != httpx.ContentTypeXIcon {
		t.Errorf("expected ContentType %v, got %v", httpx.ContentTypeXIcon, contentType)
	}
}

func TestGetFaviconHandler404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/favicon.ico", nil)
	w := httptest.NewRecorder()

	s := handler.Server{}
	s.GetFaviconHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404 OK, got %d", resp.StatusCode)
	}
}
