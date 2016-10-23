
package main
import "fmt"

func main() {
  var input float64

  fmt.Print("Enter a number: ")
  fmt.Scanf("%f", &input)

  output := input * 2 // Double the number
  fmt.Println(input, "doubled is", output)
}