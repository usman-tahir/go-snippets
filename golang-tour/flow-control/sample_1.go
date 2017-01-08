package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		fmt.Println("sum is currently:", sum)
	}

	fmt.Println("\n---------------\n")

	// Alternative for-loop with optional init and post statements
	sum = 1
	for ; sum < 10; {
		sum += sum
		fmt.Println("sum is currently:", sum)
	}

	fmt.Println("\n---------------\n")
	
	// Semicolons dropped entirely
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum is currently:", sum)
	}
}