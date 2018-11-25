package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%s' is not '%s'", got, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Frendh", func(t *testing.T) {
		got := Hello("Claire", "french")
		expected := "Bonjour, Claire"
		assertCorrectMessage(t, got, expected)
	})

}
