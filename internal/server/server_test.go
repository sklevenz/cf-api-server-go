package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Record the response
	rec := httptest.NewRecorder()

	// Call the handler directly
	handler(rec, req)

	// Check the response code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}

	// Check the response body
	expectedBody := "Hello, World!\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, rec.Body.String())
	}
}
