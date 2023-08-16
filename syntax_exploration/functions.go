package main

import "fmt"


func sub(x, y int) int {
  return x - y
}

func concat(s1 string, s2 string) string {
  return s1 + " " + s2 
}

func getNames(firstName, lastName string) (string, string) {
  return firstName, lastName
}

func main() {
  firstName := "Vitor"
  lastName := "Andrade"
  fullName := concat(firstName, lastName)
  var n1 int = 500
  var n2 int = 300
  subResult := sub(n1, n2)
  fmt.Printf("My fullname is %q\n", fullName)
  fmt.Printf("%d - %d = %d\n", n1, n2, subResult)
  fn, _ := getNames(firstName, lastName)
  fmt.Printf("Firstname is: %s\n", fn)
}
