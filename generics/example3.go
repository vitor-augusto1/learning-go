package main

import "fmt"

type CustomData interface {
  string | int | []byte | []rune
}

type User[T CustomData] struct {
  ID int
  Name string
  Data T
}

func main() {
  u := User[int]{
    ID: 1,
    Name: "Vitor",
    Data: 121293123182093,
  }
  fmt.Println(u)
}
