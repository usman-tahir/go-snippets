
package main
import "fmt"

func main() {
  index := 1
  for index <= 10 {
    fmt.Println(index)
    index += 1
  }

  /*
  // Alternative for loop
  for i := 0; i <= 10; i+= 1 {
    fmt.Println(i)
  }
  */

  // Even and odd numbers
  for i := 1; i <= 10; i += 1 {
    if (i % 2 == 0) {
      fmt.Println(i, "-> even")
    } else {
      fmt.Println(i, "-> odd")
    }
  }

  // Switch statement
  direction := ""
  fmt.Printf("Enter a cardinal direction (N, E, S, W): ")
  fmt.Scanf("%s", &direction)

  switch(direction) {
    case "N": fmt.Println("North")
    case "E": fmt.Println("East")
    case "S": fmt.Println("South")
    case "W": fmt.Println("West")
    default: fmt.Println("Unknown")
  }
}