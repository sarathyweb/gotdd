package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sarathy")
	want := "Hello, Sarathy"
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
