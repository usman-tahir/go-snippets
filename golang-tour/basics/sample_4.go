package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func new_add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 2))
	fmt.Println(new_add(2, 2)) // Also 4
}