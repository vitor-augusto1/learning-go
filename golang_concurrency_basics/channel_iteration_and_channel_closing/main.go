package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  channel := make(chan string)
  go throwingNinjaStars(channel)

  for {
    message, open := <-channel
    if !open {
      break
    }
    fmt.Println(message)
  }
}

func throwingNinjaStars(channel chan string) {
  rand.Seed(time.Now().UnixNano())
  numRounds := 3
  for i := 0; i < numRounds; i++ {
    score := rand.Intn(10)
    channel <- fmt.Sprintf("You scored: %d", score)
  }
  close(channel)
}
