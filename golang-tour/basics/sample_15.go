package main

import (
	"fmt"
)

const (
	/*
		Create a large number by shifting 1 bit left 100 places (The binary number
		1 followed by 100 zeroes)
	*/
	Big = 1<<100

	// Shift the constant Big to the right 99 places (to end up with 1<<1 = 2)
	Small = Big>>99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}