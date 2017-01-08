package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(" world!") // Executes after the program reaches its end point
	fmt.Print("Hello")
}