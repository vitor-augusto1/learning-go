package main

import (
  "fmt"
  "time"
)

// We can use channels to implement sharing memory by communicating.

// Go provides a unique concurrency synchronization technique, channel.
// Channels make goroutines share memory by communicating. We can view a
// channel as an internal FIFO (first in, first out) queue within a program.
// Some goroutines send values to the queue (the channel) and some other
// goroutines receive values from the queue.

// Along with transferring values (through channels), the ownership of some
// values may also be transferred between goroutines. When a goroutine sends a
// value to a channel, we can view the goroutine releases the ownership of some
// values (which could be accessed through the sent value). When a goroutine
// receives a value from a channel, we can view the goroutine acquires the
// ownership of some values (which could be accessed through the received value)

// Although Go also supports traditional concurrency synchronization techniques.
// only channel is first-class citizen in Go. Channel is one kind of types in Go,
// so we can use channels without importing any packages. On the other hand,
// those traditional concurrency synchronization techniques are provided in the
// sync and sync/atomic standard packages.

// The zero values of channel types are represented with the predeclared
// identifier nil. A non-nil channel value must be created by using the built-in
// make function. For example, make(chan int, 10) will create a channel whose 
// element type is int. The second argument of the make function call specifies
// the capacity of the new created channel. The second parameter is optional
// and its default value is zero.

// bidirectional channel == chan T
// send-only channel type == chan<- T
// receive-only channel type == <-chan T

func main() {
  c := make(chan int) // An unbuffered channel

  go func(ch chan<- int, x int) {
    fmt.Println("Go routine 1...")
    time.Sleep(time.Second)
    ch <- x*x
  }(c, 3)

  done := make(chan struct{})

  go func(ch <-chan int) {
    fmt.Println("Go routine 2...")
    n := <-ch
    fmt.Println(n)
    time.Sleep(time.Second)
    done <- struct{}{}
    fmt.Println("Done here...")
  }(c)

  // Making sure that all goroutines are done before ending the main func
  <-done

  fmt.Println("Bye")
}
