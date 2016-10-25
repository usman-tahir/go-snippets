
package main
import "fmt"

func first() {
  fmt.Println("first")
}

func second() {
  fmt.Println("second")
}

func main() {
  defer second() // Moves the call to the end of the function
  first()

  // Panic example
  panic("PANIC")
  s := recover()
  fmt.Println(s)
}