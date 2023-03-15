package main

import (
	"math/rand"
	"os"
)

func main() {
	random := rand.Intn(100)
	LoopUntilFound(os.Stdout, os.Stdin, random)
}
