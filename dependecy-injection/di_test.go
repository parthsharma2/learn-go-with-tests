package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Parth")

	got := buffer.String()
	want := "Hello, Parth"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
