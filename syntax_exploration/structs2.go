package main

import "fmt"


type user struct {
  name string
  number uint64
}

type messageToSend struct {
  message string
  sender user
  recipient user
}

func canSendMessage(mToSend messageToSend) bool {
  if mToSend.sender.name == "" || mToSend.recipient.name == "" {
    return false
  }
  if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
    return false
  }
  return true
}

func main() {
  newMessage := messageToSend {
    message: "Hi there!",
    sender: user{name: "Vitor", number: 148255510981},
    recipient: user{name: "Other Vitor", number: 148255510982},
  }
  fmt.Printf("Can i send the message ? %t", canSendMessage(newMessage))
}
