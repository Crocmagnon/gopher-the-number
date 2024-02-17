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
	tests := []struct {
		name       string
		random     int
		guess      int
		wantStatus string
	}{
		{"guess is lower than random", 50, 34, "Try higher."},
		{"guess is higher than random", 50, 63, "Nope, lower."},
		{"guess is equal to random", 50, 50, "Good job!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, won := CheckGuess(tt.guess, tt.random)
			if status != tt.wantStatus {
				t.Errorf("got %q want %q", status, tt.wantStatus)
			}
			if tt.guess == tt.random && !won {
				t.Error("want won")
			}
		})
	}
}

func TestLoopUntilFound(t *testing.T) {
	tests := []struct {
		name   string
		inputs []int
		random int
		want   []string
	}{
		{
			"found on first try",
			[]int{37},
			37,
			[]string{"your guess: ", "Good job!", "You got it right in 1 try."},
		},
		{
			"found on second try",
			[]int{50, 37},
			37,
			[]string{
				"your guess: ",
				"Nope, lower.",
				"your guess: ",
				"Good job!",
				"You got it right in 2 tries.",
			},
		},
		{
			"found after many tries",
			[]int{50, 25, 32, 37},
			37,
			[]string{
				"your guess: ",
				"Nope, lower.",
				"your guess: ",
				"Try higher.",
				"your guess: ",
				"Try higher.",
				"your guess: ",
				"Good job!",
				"You got it right in 4 tries.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := convertIntInputsToStrings(tt.inputs)
			inputs = append(inputs, "")
			inputBuffer := strings.NewReader(strings.Join(inputs, "\n"))

			wants := append(tt.want, "")
			want := strings.Join(wants, "\n")

			outputBuffer := bytes.Buffer{}

			LoopUntilFound(&outputBuffer, inputBuffer, tt.random)

			got := outputBuffer.String()
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func convertIntInputsToStrings(inputs []int) []string {
	inputStrings := make([]string, len(inputs)+1)
	for i, input := range inputs {
		inputStrings[i] = strconv.Itoa(input)
	}
	return inputStrings
}
