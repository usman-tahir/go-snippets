package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // Of type Vertex
	v2 = Vertex{X: 1} // Y: 0 is implicitly set
	v3 = Vertex{} // X: 0 and Y:0 are implicitly set
	p = &Vertex{1, 2} // Of type *Vertex
)

func main() {
	fmt.Println(v1, p, *p, v2, v3)
}