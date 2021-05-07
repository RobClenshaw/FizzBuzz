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

	expectedCode := http.StatusOK

	if rr.Code != expectedCode {
		t.Errorf("Ready request returned wrong status. Expected %v but got %v", expectedCode, rr.Code)
	}
}

func TestReturnsNotReadyIfPriorDataRequests(t *testing.T) {
	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Fatal(err)
	}

	count := 1
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleReady(w, r, &count)
	})

	handler.ServeHTTP(rr, req)

	expectedCode := http.StatusInternalServerError

	if rr.Code != expectedCode {
		t.Errorf("Ready request returned wrong status. Expected %v but got %v", expectedCode, rr.Code)
	}
}

func TestCountShouldIncrementOnADataRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/data/15", nil)
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleData(w, r, "myHost", "3", "Fizz", &count)
	})

	handler.ServeHTTP(rr, req)

	if count != 1 {
		t.Errorf("Expected count to increment to 1 but its value is %v", count)
	}
}
