package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i // A pointer that references i
	fmt.Println(*p) // Read the value 

	*p = 21 // Set the value of i through a reference to the pointer
	fmt.Println(i) // Read the new value of i

	p = &j // A pointer that references j
	*p = *p / 37 // Divide the value of j using the pointer
	fmt.Println(j) // Read the new value of j
}