package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello, world"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	hi := Hi("ben")
	should := "hello, ben"

	if hi != should {
		t.Errorf("got '%s' want '%s'", hi, should)
	}

}
