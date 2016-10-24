
package main
import "fmt"

func main() {
  // Create a new map
  x := make(map[string]int)
  x["a"] = 10
  fmt.Println(x["a"])

  // Add a new key to the map
  x["b"] = 20
  fmt.Println(x)

  // Delete a key from the map
  delete(x, "b")
  fmt.Println(x)

  // Nested maps (JSON lookalike)
  students := map[string]map[string]string {
    "Usman Tahir": map[string]string {
      "name": "Usman Tahir",
      "age": "21",
    },
  }
  
  fmt.Println(students)
}