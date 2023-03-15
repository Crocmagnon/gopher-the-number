package main

import (
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

func assertCheckGuess(t *testing.T, guess, random int, wantStatus string, wantWon bool) {
	status, won := CheckGuess(guess, random)
	if status != wantStatus {
		t.Errorf("got %q want %q", status, wantStatus)
	}
	if won != wantWon {
		t.Errorf("got %v want %v", won, wantWon)
	}
}
