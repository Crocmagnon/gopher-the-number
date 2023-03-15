package lib

func CheckGuess(guess, random int) (string, bool) {
	if guess == random {
		return "Good job!", true
	}
	if random > guess {
		return "Try higher.", false
	}
	return "Nope, lower.", false
}
