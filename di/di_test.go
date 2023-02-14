package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Renato")

	got := buffer.String()
	want := "Hello, Renato"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
