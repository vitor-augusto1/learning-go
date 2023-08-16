package main

import "fmt"

type messageToSend struct {
  message string
  phoneNumber uint64
}

func test(m messageToSend) {
  fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
  fmt.Println("=======================================")
}

type car struct {
  Make string
  Model string
  Height int
  Width int
  FrontWheel Wheel
  BackWheel Wheel
}

type Wheel struct {
  Radius int
  Material string
}

func main() {
  newMessage := messageToSend {
    phoneNumber: 148255510981,
    message: "Thanks for signing up",
  }
  test(newMessage)
  myCar := car{}
  myCar.BackWheel.Radius = 5
  fmt.Printf("Back wheel's radius: %v\n", myCar.BackWheel.Radius)
}
