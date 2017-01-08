package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number with an upper bound of 10
	fmt.Println("My favorite number is", rand.Intn(10))
}