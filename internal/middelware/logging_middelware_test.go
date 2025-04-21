package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware_Integration(t *testing.T) {
	wasCalled := false
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wasCalled = true
		w.WriteHeader(http.StatusTeapot) // just for fun
	})

	wrapped := LoggingMiddleware(handler)

	req := httptest.NewRequest(http.MethodPost, "/foo", nil)
	rr := httptest.NewRecorder()

	wrapped.ServeHTTP(rr, req)

	if !wasCalled {
		t.Errorf("Expected wrapped handler to be called")
	}

	if rr.Code != http.StatusTeapot {
		t.Errorf("Expected status 418, got %d", rr.Code)
	}
}
