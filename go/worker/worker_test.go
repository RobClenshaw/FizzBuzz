package main

import (
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
