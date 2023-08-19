package main

import "fmt"

func getMessageWithRetries() [3]string {
  return [3]string{
    "Click here to sign up",
    "Hey, please click here",
    "We beg you to sign up",
  }
}

func sendMessage(name string, doneAt int) {
  fmt.Printf("Sending to %v...", name)
  fmt.Println()
  
  messages := getMessageWithRetries()
  for i := 0; i < len(messages); i++ {
    msg := messages[i]
    fmt.Printf(`Sending: "%v"`, msg)
    fmt.Println()
    if i == doneAt {
      fmt.Println("They responded!")
      break
    }
    if i == len(messages)-1 {
      fmt.Println("Complete failure")
    }
  }
}

func main() {
  sendMessage("Bob", 0)
  sendMessage("John", 1)
  sendMessage("Doe", 2)
  sendMessage("Bar", 3)
}
