package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
  if len(os.Args) < 1 {
    fmt.Println("Please provide me more arguments.")
    os.Exit(1)
  }

  arguments := os.Args
  var sumOfNumbers float64
  var validNumber []float64
  for i := 0; i < len(arguments); i++ {
    currFloatNum, _ := strconv.ParseFloat(arguments[i], 64)
    if currFloatNum == 0 {
      continue
    }
    sumOfNumbers += currFloatNum
    validNumber = append(validNumber, currFloatNum)
  }
  average := sumOfNumbers / float64(len(validNumber))
  fmt.Println(">> Average >> ", average)
}
