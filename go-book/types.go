
package main
import "fmt"

func main() {
  // Numbers
  fmt.Println("1 + 1 =", (1 + 1))
  fmt.Println("1 + 1 =", (1 + 1))

  // Strings
  fmt.Println(len("Hello World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello" + " World")

  // Booleans
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
}
