package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}
	t.Run("Saying Hello to People", func(t *testing.T) {
		got := Hello("Sarathy", "")
		want := "Hello, Sarathy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello world when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spainish", func(t *testing.T) {
		got := Hello("Sarathy", "Spainish")
		want := "Hola, Sarathy"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Sarathy", "French")
		want := "Bonjour, Sarathy"
		assertCorrectMessage(t, got, want)
	})

}
