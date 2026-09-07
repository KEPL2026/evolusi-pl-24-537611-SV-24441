package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHelloWorld(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	HandleHelloWorld(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "Hello World!"
	if rr.Body.String() != expected {
		t.Errorf("Got %v, want %v", rr.Body.String(), expected)
	}
}

func TestHandle404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/afif", nil)
	rr := httptest.NewRecorder()

	Handle404(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Got %d, want %d", rr.Code, http.StatusNotFound)
	}
}
