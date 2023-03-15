package main

import (
	"bytes"
	"strconv"
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
		inputs := []int{37}
		random := 37
		want := []string{
			"your guess: ",
			"Good job!",
			"You got it right in 1 try.",
		}

		assertLoopUntilFound(t, inputs, random, want)
	})
	t.Run("found on second try", func(t *testing.T) {
		inputs := []int{50, 37}
		random := 37
		want := []string{
			"your guess: ",
			"Nope, lower.",
			"your guess: ",
			"Good job!",
			"You got it right in 2 tries.",
		}

		assertLoopUntilFound(t, inputs, random, want)
	})
	t.Run("found after many tries", func(t *testing.T) {
		inputs := []int{50, 25, 32, 37}
		random := 37
		want := []string{
			"your guess: ",
			"Nope, lower.",
			"your guess: ",
			"Try higher.",
			"your guess: ",
			"Try higher.",
			"your guess: ",
			"Good job!",
			"You got it right in 4 tries.",
		}

		assertLoopUntilFound(t, inputs, random, want)
	})
}

func assertLoopUntilFound(t testing.TB, intInputs []int, random int, outputs []string) {
	t.Helper()

	inputs := convertIntInputsToStrings(intInputs)
	inputs = append(inputs, "")
	inputBuffer := strings.NewReader(strings.Join(inputs, "\n"))

	outputs = append(outputs, "")
	want := strings.Join(outputs, "\n")

	outputBuffer := bytes.Buffer{}

	LoopUntilFound(&outputBuffer, inputBuffer, random)

	got := outputBuffer.String()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func convertIntInputsToStrings(inputs []int) []string {
	inputStrings := make([]string, len(inputs)+1)
	for i, input := range inputs {
		inputStrings[i] = strconv.Itoa(input)
	}
	return inputStrings
}
