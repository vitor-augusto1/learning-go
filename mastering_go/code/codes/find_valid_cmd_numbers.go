package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
  if len(os.Args) < 1 {
    fmt.Println("Please provide more arguments.")
    os.Exit(1)
  }

  arguments := os.Args
  var sum int64
  for i := 0; i < len(arguments); i++ {
    currNum, _ := strconv.ParseInt(os.Args[i], 10, 64)
    if currNum == 0 {
      continue
    }
    sum += currNum
  }
  fmt.Println(">> Sum of valid numbers >> ", sum)
}
