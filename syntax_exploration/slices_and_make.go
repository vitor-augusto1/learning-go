package main

import "fmt"

func main() {
  randomNumbers := [10]int{1, 22, 51, 129, 73, 00, 65, 49, 28, 67}
  for i := 0; i < len(randomNumbers); i++ {
    fmt.Println(randomNumbers[i])
  }
  slice1 := randomNumbers[2:]
  fmt.Println(slice1)
  slice2 := randomNumbers[:8]
  fmt.Println(slice2)
  slice3 := randomNumbers[:]
  fmt.Println(slice3)
  myNewSlice := make([]int, len(randomNumbers), 20)
  fmt.Println(myNewSlice[:])
  sliceLiteral := []string{"I", "love", "go"}
  fmt.Println(sliceLiteral[:])
  a := make([]int, 3)
  a[0] = 1
  a[1] = 10
  a = append(a, 100)
  fmt.Println(a)
}
