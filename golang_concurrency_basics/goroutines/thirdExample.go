package main

import "fmt"

func main() {
  myChannel := make(chan string)

  // Fork
  go func () {
    // Sender
    myChannel <- "DATA HERE"
  }()

  // Join point
  msg := <-myChannel // Receiver
  fmt.Println("MSG: ", msg)
}
