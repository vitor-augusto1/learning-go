package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
  return n%2 == 0
}
// Just waiting for the goroutines to finish before exiting

func main() {
  n := 0
  var m sync.Mutex

  // Now, both goroutines call m.Lock() before accessing `n`
  // and call m.Unlock once they are done
  go func () {
    m.Lock()
    defer m.Unlock()

    nIsEven := isEven(n)
    time.Sleep(5 * time.Millisecond)
    if nIsEven {
      fmt.Println(n, " is even...")
      return
    }
    fmt.Println(n, " is odd...")
  }()

  go func() {
    m.Lock()
    n++ 
    m.Unlock()
  }()

  time.Sleep(time.Second)
}

