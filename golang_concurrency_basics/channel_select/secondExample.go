package main

import "fmt"

func main() {
  myChannel := make(chan string)
  anotherChannel := make(chan string)

  go func() {
    myChannel <- "DATA from myChannel"
  }()

  go func() {
    anotherChannel <- "DATA from anotherChannel"
  }()

  select {
  case msgFromMyChannel := <- myChannel:
    fmt.Println(msgFromMyChannel)
  case msgFromAnotherChannel := <- anotherChannel:
    fmt.Println(msgFromAnotherChannel)
  }
}
