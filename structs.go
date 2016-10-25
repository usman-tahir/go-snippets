
package main
import "fmt"
import "math"

// Rectangle struct
type Rectangle struct {
  x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt((a * a) + (b * b))
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return (l * w)
}

type Circle struct {
  x float64
  y float64
  r float64
}

func circleArea(c *Circle) float64 {
  return math.Pi * c.r * c.r
}

// Alternative function for invocation
func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (c *Circle) diameter() float64 {
  return 2 * math.Pi * c.r
}

func main() {
  // c := new(Circle)
  // fmt.Println(c) // Currently defaults to &{0, 0, 0}

  // Initialize a circle with some other values
  c := Circle {x: 1, y: 2, r: 3}
  fmt.Println(c)
  // c := Circle {1, 2, 3}
  // c := &Circle{1, 2, 3}
  // fmt.Println(circleArea(&c))
  fmt.Println(c.area())
  fmt.Println(c.diameter())

  r := Rectangle {
    0, 0, 10, 10,
  }
  fmt.Println(r.area())

}