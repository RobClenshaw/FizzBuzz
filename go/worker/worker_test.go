package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsStringWhenMultipleOfDivisor(t *testing.T) {
	output := getOutputString(6, 3, "Fizz")
	if output != "Fizz" {
		t.Errorf("Expected \"Fizz\" for 3 but got \"%s\"", output)
	}
}

func TestReturnsEmptyStringWhenNotMultipleOfDivisor(t *testing.T) {
	output := getOutputString(7, 3, "Fizz")
	if output != "" {
		t.Errorf("Expected empty string for 7 but got \"%s\"", output)
	}
}

func TestReturnsReadyIfNoDataRequests(t *testing.T) {
	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleReady(w, r, &count)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Ready request returned wrong status. Expected %v but got %v", http.StatusOK, rr.Code)
	}
}
