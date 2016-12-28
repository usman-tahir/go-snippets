
package main
import "fmt"

func main() {
  // Basic for-loop with iteration inside the loop
  a := 0
  for a < 10 {
    fmt.Println(a)
    a += 1
  }

  // "Full" for-loop with iteration outside the loop
  for b := 0; b < 10; b += 1 {
    fmt.Println(b)
  }

  // Conditional printing inside the for-loop
  for c := 1; c <= 10; c += 1 {
    if c % 2 == 0 {
      fmt.Println(c, "is even")
    } else {
      fmt.Println(c, "is odd")
    }
  }

  // Using switch
  var d int
  fmt.Println("Enter '1' or '2': ")
  fmt.Scanf("%d", &d)
  switch d {
  case 1:
    fmt.Println("You entered '1'")
  case 2:
    fmt.Println("You entered '2'")
  default:
    fmt.Println("You did not enter '1' or '2'")
  }
}
