
package main
import "fmt"

func main() {
  var hello string = "Hello World!"
  fmt.Println(hello)
  
  // Change the variable's value
  hello = "Goodbye World!"
  fmt.Println(hello)

  // String comparison
  var first string = "foo"
  var second string = "bar"
  fmt.Println(first == second) // false
  second = "foo"
  fmt.Println(first == second) // true

  // Type inference
  third := "baz"
  fmt.Println(third)

  // Shorthand/minimalist assignment
  var (
    a = 1
    b = 2
    c = 3
  )
  fmt.Println(a, b, c)

}