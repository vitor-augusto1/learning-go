package main

import "fmt"

type Num interface {
  int | int64 | float64 | float32
}

func add[T Num](a T, b T) T {
  return a + b
}

func main() {
  sum := add(1.2, 2.0)
  fmt.Println("Sum: ", sum)
}

func mapValues(values []int, mapFunc func(n int) int) []int {
  var newValues []int
  for _, value := range values {
    newValue := mapFunc(value)
    newValues := append(newValues, newValue)
  }
  return newValues
}

func main() {
  mapResult := mapValues([]int{1, 2, 3}, func(n int) int {
    return n * 5
  })
  fmt.Printf("Result: %+v\n", mapResut)
}
