package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	max := 100
	fmt.Printf("I've picked a number between 0 and %d. Can you guess it?\n", max)
	random := rand.Intn(max)
	LoopUntilFound(os.Stdout, os.Stdin, random)
}
