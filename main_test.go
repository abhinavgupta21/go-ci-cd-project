package main

import (
	"testing"
)

func TestGetGreeting(t *testing.T) {
	expected := "Hello, CI/CD with Github Actions Project!"
	if actual := getGreeting(); actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
