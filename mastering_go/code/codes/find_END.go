package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  var f *os.File
  f = os.Stdin
  defer f.Close()

  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    currStr := scanner.Text()
    fmt.Println(">> Currently string >> ", currStr)
    if currStr == "END" {
      fmt.Println("END string found. Exiting...")
      os.Exit(0)
    }
  }
}
