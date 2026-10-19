package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogResponse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	newHandler().ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", response.Code)
	}
	if response.Header().Get("X-Service") != "catalog" {
		t.Fatalf("X-Service = %q, want catalog", response.Header().Get("X-Service"))
	}
	if response.Body.String() != "hello from catalog" {
		t.Fatalf("body = %q, want hello from catalog", response.Body.String())
	}
}

func TestUnknownRoute(t *testing.T) {
	response := httptest.NewRecorder()
	newHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/unknown", nil))
	if response.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", response.Code)
	}
}
