package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Parth", "")
		want := "Hello, Parth"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Parth", "Spanish")
		want := "Hola, Parth"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Parth", "French")
		want := "Bonjour, Parth"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Parth", "Hindi")
		want := "Namaste, Parth"
		assertCorrectMessage(t, got, want)
	})
}
