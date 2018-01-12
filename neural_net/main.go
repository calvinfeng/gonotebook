package main

import (
	"math/rand"
	"time"
)

func main() {
	// If seed is not set, we will expect to get the result random numbers in runtime every time we run the program. This
	// particular seed will give the program pseudo-randomness.
	rand.Seed(time.Now().UTC().Unix())
}
