package main

import (
  "fmt"
)

type Person struct {
  name string
  age int
}

func older(personOne, personTwo Person) (Person, int) {
  if personOne.age > personTwo.age {
    return personOne, personOne.age - personTwo.age
  }
  return personTwo, personTwo.age - personOne.age
}

func main() {
  var p1 Person
  p1.name, p1.age = "Tom", 18

  p2 := Person{"Bob", 22}

  whoIsOlder, ageDifference := older(p1, p2)
  fmt.Printf("Of %s and %s, %s is older by %d years.\n", p1.name, p2.name, whoIsOlder.name, ageDifference)
}
