package main

import (
  "fmt"
  "example.com/packages/utils"
)

func main() {
  greeting := fmt.Sprintf("Hello, %s", "Vitor")
  fmt.Println(greeting)
  fmt.Println(utils.stringLen("Test"))
}
