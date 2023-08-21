package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
  now := time.Now()
  userId := 10
  // Creating a new chanel
  responseChanel := make(chan string, 128)
  wg := &sync.WaitGroup{}

  go fetchUserData(userId, responseChanel, wg)
  go fetchUserRecommendations(userId, responseChanel, wg)
  go fetchUserLikes(userId, responseChanel, wg)

  wg.Add(3)
  wg.Wait()

  close(responseChanel)

  for resp := range responseChanel {
    fmt.Println(resp)
  }

  fmt.Println(time.Since(now))
}


func fetchUserData(userId int, responseChanel chan string, wg *sync.WaitGroup) {
  time.Sleep(80 * time.Millisecond)
  responseChanel <- "User data"

  wg.Done()
}

func fetchUserRecommendations(userId int, responseChanel chan string, wg *sync.WaitGroup) {
  time.Sleep(120 * time.Millisecond)
  responseChanel <- "User recommendations"

  wg.Done()
}

func fetchUserLikes(userId int, responseChanel chan string, wg *sync.WaitGroup) {
  time.Sleep(50 * time.Millisecond)
  responseChanel <- "User likes"

  wg.Done()
}
