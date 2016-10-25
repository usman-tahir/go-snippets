
package main
import "fmt"

type Person struct {
  name string
}

type Android struct {
  Person
  model string
}

func (p *Person) introduce() {
  fmt.Println("Hello, my name is", p.name + ".")
}

// Interface example
type Shape interface {
  area() float64
}

// Composition of multiple smaller shapes
type MultiShape struct {
  shapes []Shape
}

func main() {
  p := Person{name: "Usman Tahir"}
  p.introduce()

  j := Person{name: "John Doe"}
  j.introduce()

  a := Android{model: "1001"}
  a.Person.name = "Android 1000"
  a.introduce()
}