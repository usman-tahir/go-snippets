
package main
import "fmt"

var g string = "Global variable"

func f() {
  fmt.Println("invocation from function f, g =", (g))
}

func main() {

  // Example of a constant
  const hello_world string = "Hello World"
  fmt.Println(hello_world)

  var x string = "Hello World"
  fmt.Println(x)

  // Alternatively, declare and initialize on two lines
  var y string
  y = "Hello World"
  fmt.Println(y)

  // Changing the value of variables
  var z string
  z = "first"
  fmt.Println(z)
  z = "second"
  fmt.Println(z)

  // String equality
  x = "hello"
  y = "world"
  fmt.Println(x == y) // false

  // Shorthand variable naming
  name := "Usman"
  fmt.Println("His name is", (name))

  // Scoping example
  fmt.Println("Invocation from main, g =", (g))
  f()

  // Defining multiple variables
  var (
    i = 5
    j = 10
    k = 15
  )

  fmt.Println(i, j, k)
}
