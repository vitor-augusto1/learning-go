package main

import (
  "fmt"
)

func main() {
  aSlice := []int{1, 2, 3, 4, 5}
  fmt.Println(aSlice)
  integers := make([]int, 2)
  fmt.Println(integers)
  integers = nil
  fmt.Println(integers)

  anArray := [5]int{-1, -2, -3, -4, -5}
  // Creating a slice from an array using the [:] notation
  refAnArray := anArray[:]
  fmt.Println(anArray)
  fmt.Println(refAnArray)
  anArray[4] = -100
  fmt.Println(refAnArray)

  s := make([]byte, 5)
  fmt.Println(s)
  twoD := make([][]int, 3)
  fmt.Println(twoD)
  fmt.Println()

  for i := 0; i < len(twoD); i++ {
    for j := 0; j < 2; j++ {
      twoD[i] = append(twoD[i], i*j)
    }
  }
}
