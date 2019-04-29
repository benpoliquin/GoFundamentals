package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ben")

	got := buffer.String()
	want := "Hello, Ben"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
