
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Checks if the string contains a substring/pattern
  fmt.Println(strings.Contains("hello", "ello")) // true

  // Counts the number of occurrences of a certain pattern
  fmt.Println(strings.Count("test", "t")) // 2
  fmt.Println(strings.Count("test", "te")) // 1

  // Checks to see if the given string starts with a certain pattern
  fmt.Println(strings.HasPrefix("hello", "he")) // true
  fmt.Println(strings.HasPrefix("hello", "lo")) // false

  // Checks to see if the given string ends with a certain pattern
  fmt.Println(strings.HasSuffix("hello", "lo")) // true
  fmt.Println(strings.HasSuffix("hello", "he")) // false

  // Returns the index of the pattern in the given string
  fmt.Println(strings.Index("hello", "e")) // 1
  fmt.Println(strings.Index("hello", "q")) // -1

  // Take a list of strings and join them together with a delimeter
  s := []string {
    "hello",
    "world",
  }
  fmt.Println(strings.Join(s, "-")) // hello-world

  // Repeat a string a given amount of times
  fmt.Println(strings.Repeat("a", 5)) // aaaaa

  // Replace a smaller string within a bigger string
  // NOTE: The number argument determines how many elements are replaced
  fmt.Println(strings.Replace("aaaaaaaaaa", "a", "b", 5)) // bbbbbaaaaa
  fmt.Println(strings.Replace("aabaabaab", "a", "c", 3)) // ccbcabaab

  // Split a string into a list of elements
  h := "Hello World!"
  fmt.Println(strings.Split(h, "")) // [H e l l o  W o r l d !]

  // Convert a string into all lowercase or uppercase
  fmt.Println(strings.ToLower("HELLO")) // hello
  fmt.Println(strings.ToUpper("hello")) // HELLO

}