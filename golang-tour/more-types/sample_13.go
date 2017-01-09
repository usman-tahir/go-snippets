package main

import (
  "fmt"
)

func printSlice(s []int) {
  fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}

func main() {
  var s []int
  printSlice(s)

  // Append on the nil slice
  s = append(s, 0)
  printSlice(s)

  // Appending to the slice allows it to grow in size
  s = append(s, 1)
  printSlice(s)

  // Appending more than one element at a time is supported
  s = append(s, 2, 3, 4)
  printSlice(s)

}