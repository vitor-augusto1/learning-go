package main

import "fmt"

func sum(nums ...int) int {
  var num int
  for i := 0; i < len(nums); i++ {
    num += nums[i]
  }
  return num
}

func main() {
  nums := []int{1, 2, 3, 5}
  total := sum(nums...)
  fmt.Println(total)
}
