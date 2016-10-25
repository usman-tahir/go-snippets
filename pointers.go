
package main
import "fmt"

func zero(xPointer *int) {
  *xPointer = 0
}

func one(xPointer *int) {
  *xPointer = 1
}

func main() {
  x := 5
  fmt.Println(x) // Before the function call (5)
  zero(&x)
  fmt.Println(x) // x has been changed to zero

  xPointer := new(int)
  one(xPointer)
  fmt.Println(*xPointer) // x is now one
}