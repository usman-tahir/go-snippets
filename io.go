
package main
import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("sample.txt")
  if (err != nil) {
    // Handle the error
    return
  }
  defer file.Close()

  // Get the file size
  stat, err := file.Stat()
  if (err != nil) {
    // Handle the error
    return
  }

  // Read the file
  bytes := make([]byte, stat.Size())
  _, err = file.Read(bytes) // Read the lines from the file
  if (err != nil) {
    // Handle the error
    return
  }

  str := string(bytes)
  fmt.Println(str) // Print out the file contents
}