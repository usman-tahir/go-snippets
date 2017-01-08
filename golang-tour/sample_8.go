package main

import (
	"fmt"
)

var i, j int = 1, 2 // i = 1, j = 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java) // 1 2 true false no!
}