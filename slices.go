
package main
import "fmt"

func main() {
  x := [5]float64 {
    1,
    2,
    3,
    4,
    5,
  }

  // Initialize the slice
  slice := x[0:2] // The first two elements
  fmt.Println(slice)

  // Appending slices
  a := []int {
    1,
    2,
    3,
  }

  // [1, 2, 3, 4, 5]
  b := append(a, 4, 5)
  fmt.Println(a, b)

  // Copying slices
  c := []int {
    1,
    2,
    3,
  }

  d := make([]int, 2)
  copy(d, c) // Copy into slice 'd' the contents of slice 'c'
  fmt.Println(c, d) // 'd' has only [1, 2] since its space limitation is 2

}