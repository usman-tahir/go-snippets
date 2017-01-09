package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("length = %d, cap = %d: %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	printSlice(s)

	// Slice the current slice (s) to give it a length of 0
	s = s[:0]
	printSlice(s)

	// Extend the length of the slice
	s = s[:4]
	printSlice(s)

	// Drop the first two values of the slice
	s = s[2:]
	printSlice(s)

}
