package main

import "fmt"

// RANGE syntax: for INDEX, ELEMENT := range SLICE {}

func showFruits(fruits []string) {
  for i, fruit := range fruits {
    fmt.Println(i, fruit)
  }
}


func main() {
  fruits := []string {"Apple", "Grape", "Banana"}
  showFruits(fruits)
}
