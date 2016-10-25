
package main
import "fmt"

func average(values []float64) float64 {
  var sum float64 = 0.00
  for i := 0; i < len(values); i += 1 {
    sum += values[i]
  }
  sum = sum / float64(len(values))
  return sum
}

// Recursive factorial
func factorial(number int) int {
  if (number == 0 || number == 1) {
    return 1
  } else {
    return number * factorial(number - 1)
  }
}

// Variadic function
func add(args ...int) int {
  total := 0
  for _, v := range(args) {
    total += v
  }
  return total
}

// Even number generator
func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func main() {
  x := []float64 {
    1, 2, 3,
  }

  fmt.Println(average(x))
  fmt.Println(factorial(4)) // 24
  fmt.Println(add(1, 2, 3)) // 6

  // Closure
  y := 0
  increment := func() {
    y++
  }
  increment() // y = 1
  fmt.Println(y)

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2

}