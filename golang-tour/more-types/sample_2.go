package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// Changing the value of the Vertex struct instance
	v.X = 4
	fmt.Println(v)

	// Using pointers with structs
	p := &v
	fmt.Println(*p) // Read the value referenced by the pointer

	// Change the value of the Vertex's field through a pointer
	p.X = 1
	fmt.Println(v)
}