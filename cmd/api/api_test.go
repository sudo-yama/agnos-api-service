package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ✅ test /health/live
func TestLiveHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	w := httptest.NewRecorder()

	liveHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

// ✅ test /health/ready
func TestReadyHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health/ready", nil)
	w := httptest.NewRecorder()

	readyHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
