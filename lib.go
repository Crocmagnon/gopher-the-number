package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func CheckGuess(guess, random int) (string, bool) {
	if guess == random {
		return "Good job!", true
	}
	if random > guess {
		return "Try higher.", false
	}
	return "Nope, lower.", false
}

func LoopUntilFound(writer io.Writer, reader io.Reader, random int) {
	scanner := bufio.NewScanner(reader)
	win := false
	count := 0
	for !win {
		fmt.Fprint(writer, "your guess: ")
		scanner.Scan()
		value := scanner.Text()
		guess, _ := strconv.Atoi(value)
		var msg string
		msg, win = CheckGuess(guess, random)
		fmt.Fprintf(writer, "\n%v\n", msg)
		count++
	}
	if count == 1 {
		fmt.Fprintln(writer, "You got it right in 1 try.")
	} else {
		fmt.Fprintf(writer, "You got it right in %d tries.\n", count)
	}
}
