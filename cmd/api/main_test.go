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

// ✅ test metricsMiddleware
func TestMetricsMiddleware(t *testing.T) {
	handler := http.HandlerFunc(liveHandler)
	wrapped := metricsMiddleware(handler)

	req := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	w := httptest.NewRecorder()

	wrapped.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

// ✅ test responseWriter.WriteHeader
func TestResponseWriter_WriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	rw := &responseWriter{ResponseWriter: w, statusCode: 200}

	rw.WriteHeader(201)

	if rw.statusCode != 201 {
		t.Errorf("expected 201, got %d", rw.statusCode)
	}
}
