package handler

import "testing"

func TestHello(t *testing.T) {
	result := "hello"

	if result != "hello" {
		t.Error("Expected hello")
	}
}
