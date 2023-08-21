package main

import (
  "fmt"
)


type Message struct {
  From string
  Payloads string
}


type Server struct {
  msgch chan Message
}


func (s *Server) StartAndListen() {
  for {
    msg := <- s.msgch
    fmt.Printf("Received message from: %s payload %s\n", msg.From, msg.Payloads)
  }
}


func sendMessageToServer(msgch chan Message, payload string) {
  msg := Message{
    From: "Joe",
    Payloads: payload,
  }
  msgch <- msg
}


func main() {
  s := &Server{
    msgch: make(chan Message),
  }
  
  go s.StartAndListen()

  sendMessageToServer(s.msgch, "Hello Sailor!")
}
