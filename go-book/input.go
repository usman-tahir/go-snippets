
package main
import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64

  fmt.Scanf("%f", (&input))
  output := input * 2 // Double the input
  fmt.Println(output)
}
