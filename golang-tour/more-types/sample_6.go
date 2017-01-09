package main

import (
	"fmt"
)

func main() {
	names := [4]string {"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	sliceA := names[0:2]
	sliceB := names[1:3]

	fmt.Println(sliceA, sliceB)

	sliceB[0] = "X"
	fmt.Println(sliceA, sliceB)
	fmt.Println(names)
}