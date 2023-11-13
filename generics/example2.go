package main

import (
  "fmt"
)

func mapValues[T any](values []T, mapFunc func(n T) T) []T {
  var newValues []T
  for _, v := range values {
    newValue := mapFunc(v)
    newValues = append(newValues, newValue)
  }
  return newValues
}

func main() {
  mapResult := mapValues([]int{1, 2, 3}, func(n int) int {
    return n * 5
  })
  fmt.Printf("Result: %+v\n", mapResult)
}
