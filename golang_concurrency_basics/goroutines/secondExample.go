package main

import "fmt"

func someFunc(num string) {
  fmt.Println("Num: ", num)
}

func main() { // Start

  // Fork's
  go someFunc("2")
  go someFunc("3")
  go someFunc("4")
  // Fork's

  // Join point

  fmt.Println("Hi")
}
