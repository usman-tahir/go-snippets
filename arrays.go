
package main
import "fmt"

func main() {
  // Array of integers
  var x [5]int

  // Assign the last value to 100
  x[4] = 100
  fmt.Println(x) // All other values are 0 by default

  // Get the average of a group of numbers
  sum := 0
  /*
  var y [5]int
  y[0] = 1
  y[1] = 2
  y[2] = 3
  y[3] = 4
  y[4] = 5
  */
  y := [5]int  {
    1,
    2,
    3,
    4,
    5,
  }

  for i := 0; i < 5; i += 1 {
    sum += y[i]
  }
  fmt.Println("Average:", (sum / len(y)))

}