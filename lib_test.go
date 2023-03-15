package main

import (
	"bytes"
	"strings"
	"testing"
)

/*
Requirements:
The app selects a random number between 1 and 100
The app then asks the player to guess the number.
If the guess is correct, then the player wins. The app displays
the number of guesses it took and exits.
If the guess is incorrect, the app prints whether the selected number
is higher or lower than the guess.
Then, the app prompts the player again, etc.
*/

func TestCheckGuess(t *testing.T) {
	t.Run("guess is lower than random", func(t *testing.T) {
		random := 63
		guess := 50
		wantStatus := "Try higher."

		assertCheckGuess(t, guess, random, wantStatus, false)
	})
	t.Run("guess is higher than random", func(t *testing.T) {
		random := 34
		guess := 50
		wantStatus := "Nope, lower."

		assertCheckGuess(t, guess, random, wantStatus, false)
	})
	t.Run("guess is equal to random", func(t *testing.T) {
		random := 50
		guess := 50
		wantStatus := "Good job!"

		assertCheckGuess(t, guess, random, wantStatus, true)
	})
}

func assertCheckGuess(t testing.TB, guess, random int, wantStatus string, wantWon bool) {
	t.Helper()
	status, won := CheckGuess(guess, random)
	if status != wantStatus {
		t.Errorf("got %q want %q", status, wantStatus)
	}
	if won != wantWon {
		t.Errorf("got %v want %v", won, wantWon)
	}
}

func TestLoopUntilFound(t *testing.T) {
	t.Run("found on first try", func(t *testing.T) {
		output := bytes.Buffer{}
		input := strings.NewReader("37\n")
		random := 37
		LoopUntilFound(&output, input, random)

		got := output.String()
		want := "your guess: \nGood job!\nYou got it right in 1 try.\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("found on second try", func(t *testing.T) {
		output := bytes.Buffer{}
		input := strings.NewReader("50\n37\n")
		random := 37
		LoopUntilFound(&output, input, random)

		got := output.String()
		want := "your guess: \nNope, lower.\nyour guess: \nGood job!\nYou got it right in 2 tries.\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("found after many tries", func(t *testing.T) {
		output := bytes.Buffer{}
		inputs := []string{"50", "25", "32", "37", ""}
		input := strings.NewReader(strings.Join(inputs, "\n"))
		random := 37
		LoopUntilFound(&output, input, random)

		got := output.String()
		wants := []string{
			"your guess: ",
			"Nope, lower.",
			"your guess: ",
			"Try higher.",
			"your guess: ",
			"Try higher.",
			"your guess: ",
			"Good job!",
			"You got it right in 4 tries.",
			"",
		}
		want := strings.Join(wants, "\n")

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
